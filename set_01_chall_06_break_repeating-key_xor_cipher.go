package main

import ("fmt";
        "os";
        "bytes";
        "io/ioutil";
        b64 "encoding/base64";
        "./hamming";
        "sort";
        "./xor";
)

type distances struct {
    keysize int
    normalized_distance float64
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func break_blocks(size int, full_block []byte) ([][]byte, int) {
    var blocks [][]byte
    y := 0
    x := 0
    num_blocks := (len(full_block)/size)
    for ; x < num_blocks; x++ {
        blocks = append(blocks, full_block[y:y+size])
        y += size
    }
    fmt.Println("==============")
    if (len(full_block)%size > 0) {
        blocks = append(blocks, full_block[y:])
    }
    return blocks, size
}

func transpose_blocks(size int, blocks [][]byte, totalsize int) [][]byte {
    var transposed [][]byte
    var temp []byte
    var stop bool
    stop = false
    for count, lowerb := 0, 0; stop == false; lowerb++ {
        for upperb := 0; upperb < len(blocks); upperb++{
            if (lowerb < (len(blocks[upperb]))){
                temp = append(temp, blocks[upperb][lowerb])
                count++
            } else if ( count == totalsize  ) {
                stop = true
                break
            }
        }
        if (temp != nil) {
            transposed = append(transposed, temp)
        }
        temp = nil
    }
    return transposed
}

func candidate_keys(transposed [][]byte) [][]int {
    var possible_keys [][]int
    for i := 0; i < len(transposed); i++ {
        temp := xor.Break(transposed[i], true)
        if (temp == nil) {
            fmt.Println("Invalid Keysize\n\n")
            return nil
        }
        possible_keys = append(possible_keys, temp)
    }
    fmt.Println("Keys")
    fmt.Println(possible_keys)
    return possible_keys
}

func guess_keys(ciphertext []byte, keysize int) [][]int {
    fmt.Printf("Trying keys with length %d...\n", keysize)
    fmt.Printf("Breaking into %d-byte blocks\n", keysize)
    blocks, size := break_blocks(keysize, ciphertext)
    transposed := transpose_blocks(size, blocks, len(ciphertext))
    return candidate_keys(transposed)

}

func main(){
    if len(os.Args) == 1 || len(os.Args) > 2 {
        fmt.Printf("\n\tUsage: %s <file>\n\n", os.Args[0])
        os.Exit(0)
    }
    fcontent, err := ioutil.ReadFile(os.Args[1])
    check(err)
    fcontent2 := bytes.Replace([]byte(fcontent), []byte{10}, nil, -1)
    ciphertext,_ := b64.StdEncoding.DecodeString(string(fcontent2))
    var max_keysize int
    if (len(ciphertext)/2)<40 {
        max_keysize = len(ciphertext)/2
    } else {
        max_keysize = 40
    }
    fmt.Printf("Max Keysize: %d\n", max_keysize)
    type keys struct {
        keysize int
        ndist float64
    }
    var median_ndistances []keys
    for x:=2;x<=max_keysize;x++ {
        ndist := hamming.Normalized_distance(ciphertext, x)
        median_ndistances=append(median_ndistances, keys{x, ndist})
    }
    //Sorting Normalized distances
    sort.SliceStable(median_ndistances, func(i, j int) bool {
                            return median_ndistances[i].ndist < median_ndistances[j].ndist})
    fmt.Println(median_ndistances)
    var probable_keysizes []int
    //Collect the 3 most probable keysizes only
    for tmp:=0; tmp<3; tmp++ {
        probable_keysizes=append(probable_keysizes, median_ndistances[tmp].keysize)
    }
    println("-----")
    for x := range probable_keysizes {
        probable_keys := guess_keys(ciphertext, probable_keysizes[x])
        if (probable_keys == nil) {
            //Invalid Keysize
            probable_keysizes[x] = 0
            continue
        }
        fmt.Printf("Best guess: \n")
        var bguess []byte
        for i:=0; i<len(probable_keys); i++ {
            //Retrieve the top scored keys
            bguess = append(bguess, byte(probable_keys[i][0]))
        }
        fmt.Println(probable_keysizes)
        fmt.Println(bguess)
        result := xor.Decipher(bguess, ciphertext)
        fmt.Printf("\n\n%s\n", result)
    }
    fmt.Println("Most probable keysizes:")
    fmt.Println(probable_keysizes)
}




