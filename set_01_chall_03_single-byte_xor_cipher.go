package main

import ("fmt"; "os"; "encoding/hex"; "sort"; "./freq_score"; "strings")

type ciphertext struct {
    ctext []byte
    key, score  int
}

func main(){
    if len(os.Args) == 1 || len(os.Args) > 3 {
        fmt.Printf("\n\tUsage: xor_decoder.go <hex string> [--pconly]\n")
        fmt.Printf("\n\t\t--pconly\tPrintable characters only\n\n")
        os.Exit(0)
    }
    in := os.Args[1]
    in_hexdecoded, _ := hex.DecodeString(in)
    only_printable := true
    ct_list := []ciphertext{}
    for x:=0;x<256;x++{
        var not_printable_char bool
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
        ct_list = append(ct_list, ciphertext{dst, key, freq_score.Score(dst)})

    }
    sort.SliceStable(ct_list, func(i, j int) bool { return ct_list[i].score > ct_list[j].score })
    if len(ct_list) > 0 {
        var limit int
        if len(ct_list) > 4 {
            //limit = len(ct_list)
            limit = 5
            fmt.Printf("Limit: %d\n", limit)
        } else {
            limit = len(ct_list)
            fmt.Printf("Limit: %d\n", limit)
        }
        for i:=0;i<limit;i++{
            fmt.Printf("Ciphertext: %s - Key: %d - Score: %d\n", strings.Trim(string(ct_list[i].ctext), "\n"), ct_list[i].key, ct_list[i].score)
            //fmt.Printf("%d\n", ct_list[i].ctext)
        }
    }
}

