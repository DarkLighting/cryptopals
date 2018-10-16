package xor

import ("fmt";
        "sort";
        "../freq_score";
        "strings"
        )

type ciphertext struct {
    ctext []byte
    key, score  int
}

func Xor(in1 byte, in2 byte) byte {
    dst := byte(int(in1) ^ int(in2))
    return dst
}

func Xor_rune(in1 rune, in2 rune) rune {
    var dst rune
    dst = rune(int(in1) ^ int(in2))
    return dst
}

func is_printable(ch byte) bool {
    if int(ch) == 10 || int(ch) == 13 || (int(ch) > 31 && int(ch) < 128) {
        return  true
    } else {
        return false
    }
}


func Break(ciphered []byte, only_printable bool) []int {
    ct_list := []ciphertext{}
    for x:=0;x<256;x++{
        var not_printable_char bool
        key := int(x)
        dst := make([]byte, len(ciphered))
        for i:=0;i<len(ciphered);i++ {
            dst[i] = byte(int(ciphered[i]) ^ key)
            if only_printable {
                if is_printable(dst[i]) {
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
            limit = 5
        } else {
            limit = len(ct_list)
        }
        var possible_keys []int
        for i:=0;i<limit;i++{
            fmt.Printf("\nCiphertext: %s  - Key: %d - Score: %d\n", strings.Trim(string(ct_list[i].ctext), "\n"), ct_list[i].key, ct_list[i].score)
            possible_keys = append(possible_keys, ct_list[i].key)
        }
        return possible_keys
    }
    return nil
}

func Decipher(candidate []byte, ciphertext []byte) []byte {
    fmt.Println("Candidate Key")
    fmt.Println(candidate)
    var plaintext []byte
    for i:=0; i < len(ciphertext); i++ {
        plaintext = append(plaintext, Xor(ciphertext[i], candidate[i%29]))
    }
    return plaintext
}

