package aesecb

import ("crypto/aes";
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func AesEcb_decrypt(ciphertext []byte, key []byte) []byte {
    block, err := aes.NewCipher(key)
    check(err)
    keysize := len(key)
    plaintext := make([]byte, len(ciphertext))
    for count:=0; count < len(ciphertext); count+=keysize {
        ct_temp := ciphertext[count:count+keysize]
        pt_temp := make([]byte, len(ct_temp))
        block.Decrypt(pt_temp, ct_temp)
        for i:=0; i < keysize; i++ {
            plaintext[count+i] = pt_temp[i]
        }
    }
    return plaintext
}
