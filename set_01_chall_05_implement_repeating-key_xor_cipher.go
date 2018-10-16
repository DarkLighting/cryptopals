package main

import ("fmt";
        "os";
        "encoding/hex";
        "io/ioutil")

func check_error(e error){
    if e!= nil {
        panic(e)
    }
}

func main(){
    file := os.Args[1]
    in, _ := ioutil.ReadFile(file)
    key := "ICE"
    key_byte := []byte(key)
    dst := make([]byte, len(in))
    for i:=0;i<len(in);i++ {
        dst[i] = byte(int(in[i]) ^ int(key_byte[i%len(key_byte)]))
    }
    fmt.Println(hex.EncodeToString(dst))
}
