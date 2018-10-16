package main

import ("fmt"; "os"; "encoding/hex"; "./xor")

func main(){
    in := os.Args[1]
    in_hexdecoded, _ := hex.DecodeString(in)
    fmt.Println(in_hexdecoded)
    fmt.Println([]byte(in))
    operand := "686974207468652062756c6c277320657965"
    operand_hexdecoded, _ := hex.DecodeString(operand)
    fmt.Println(operand_hexdecoded)
    dst := make([]byte, len(in_hexdecoded))
    if len(in_hexdecoded) == len(operand_hexdecoded) {
        dst = xor.Xor(in_hexdecoded, operand_hexdecoded)
        fmt.Println(hex.EncodeToString(dst))
    }
}
