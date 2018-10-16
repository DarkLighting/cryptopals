package main

import ("fmt";
        "os";
        "bytes"
        "io/ioutil";
        b64 "encoding/base64";
        "./aesecb"
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
    fcontent2 := bytes.Replace([]byte(fcontent), []byte{10}, nil, -1)
    ciphertext,_ := b64.StdEncoding.DecodeString(string(fcontent2))
    fmt.Printf("\nTrying to decrypt ciphertext: \n\n")
    key := []byte("YELLOW SUBMARINE")
    plaintext := aesecb.AesEcb_decrypt(ciphertext, key)
    fmt.Printf("%s\n\n", plaintext)
}
