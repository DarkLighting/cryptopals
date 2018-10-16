package main

import ("fmt";
        "os";
        "bytes"
        "io/ioutil";
        "strings";
        "./hamming";
        "sort";
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main(){
    if len(os.Args) == 1 || len(os.Args) > 2 {
        fmt.Printf("\n\tUsage: %s <file>\n\n", os.Args[0])
        os.Exit(0)
    }
    fcontent, err := ioutil.ReadFile(os.Args[1])
    check(err)
    fcontent = bytes.TrimSuffix(fcontent, []byte{10})
    lines := strings.Split(string(fcontent), "\n")
    var linebytes [][]byte
    for x:=0;x<len(lines);x++{
        l:=[]byte(lines[x])
        fmt.Println(l)
        linebytes = append(linebytes, l)
    }
    type keys struct {
        keysize int
        line int
        ndist float64
    }
    keysizes := []int{16, 24, 32}
    var median_ndistances []keys
    for x:=0;x<len(keysizes);x++ {
        for i:=0; i<len(linebytes); i++ {
            ndist := hamming.Normalized_distance(linebytes[i], keysizes[x])
            median_ndistances=append(median_ndistances, keys{keysizes[x], i, ndist})
        }
    }
    sort.SliceStable(median_ndistances, func(i, j int) bool {
                            return median_ndistances[i].ndist < median_ndistances[j].ndist})
    fmt.Println(median_ndistances)
    fmt.Printf("\n\nMost probable AES-ECB string:\n")
    fmt.Printf("Line %d\n", median_ndistances[0].line+1)
    fmt.Println(linebytes[median_ndistances[0].line])
    fmt.Printf("\n\n5 lowest Hamming distances:\n")
    fmt.Printf("\n\nKeysize / Line Index / Hamming Distance\n\n")
    for x:=0;x<10;x++ {
        fmt.Println(median_ndistances[x])
    }
}

