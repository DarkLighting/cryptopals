package main

import ("fmt"; "os"; b64 "encoding/base64"; "encoding/hex")

func main(){
    hexin := os.Args[1]
    data, _ := hex.DecodeString(hexin)
    fmt.Println(data)
    dataout := b64.StdEncoding.EncodeToString([]byte(data))
    fmt.Println(dataout)
}
