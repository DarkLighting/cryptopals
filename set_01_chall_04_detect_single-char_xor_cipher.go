package main

import ("fmt";
        "os";
        "encoding/hex";
        "sort";
        "./freq_score";
        "io/ioutil";
        "strings")

type ciphertext struct {
    ctext []byte
    key, score, line_number  int
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main(){
    if len(os.Args) == 1 || len(os.Args) > 2 {
        fmt.Printf("\n\tUsage: xor_decoder.go <file>\n\n")
        os.Exit(0)
    }
    fcontent, err := ioutil.ReadFile(os.Args[1])
    check(err)
    lines := strings.Split(string(fcontent), "\n")
    for x:=0;x<len(lines);x++{
        try_decode(lines[x], x+1)
    }
}

func try_decode(in string, line_number int){
    in_hexdecoded, _ := hex.DecodeString(in)
    var not_printable_char bool
    only_printable := true
    ct_list := []ciphertext{}
    for x:=0;x<256;x++{
        key := int(x)
        dst := make([]byte, len(in_hexdecoded))
        for i:=0;i<len(in_hexdecoded);i++ {
            dst[i] = byte(int(in_hexdecoded[i]) ^ key)
            if only_printable {
                if int(dst[i]) == 10 || (int(dst[i]) > 31 && int(dst[i]) < 127) {
                    not_printable_char = false
                } else {
                    not_printable_char = true
                    break
                }
            }
        }
        if only_printable {
            if not_printable_char { continue }
        }
        ct_list = append(ct_list,
                ciphertext{dst, key, freq_score.Score(dst), line_number})

    }
    sort.SliceStable(ct_list, func(i, j int) bool {
            return ct_list[i].score > ct_list[j].score })
    if len(ct_list) > 0 {
        fmt.Printf("\nLine Number: %d\n", line_number)
        var limit int
        if len(ct_list) > 4 {
            limit = 5
        } else {
            limit = len(ct_list)
        }
        for i:=0;i<limit;i++{
            fmt.Printf("Ciphertext: %s - Key: %d - Score: %d\n",
                    strings.Trim(string(ct_list[i].ctext), "\n"),
                    ct_list[i].key,
                    ct_list[i].score)
        }
    }
}
