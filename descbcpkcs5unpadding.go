package main

import (
	"crypto/cipher"
	"crypto/des"
	"encoding/base64"
	"fmt"
)

type T struct {
	A int
	B string
}

func main() {

	shops := make([]int64, 0)

	fmt.Printf("%#v--%d \n", shops, len(shops))

	mobile := "9FJ4HeLJOjZzGU58/UlmfQ\u003d\u003d"
	mobiledec64, _ := base64.StdEncoding.DecodeString(mobile)
	iv := []byte{1, 2, 3, 4, 5, 6, 7, 8}
	ret, _ := DesDecrypt(mobiledec64, []byte("JE9DvDv5"), iv)
	fmt.Printf("%#v,%s", ret, string(ret))

}

func DesDecrypt(crypted, key, iv []byte) ([]byte, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}

	blockMode := cipher.NewCBCDecrypter(block, iv)
	origData := make([]byte, len(crypted))
	// origData := crypted
	blockMode.CryptBlocks(origData, crypted)
	origData = PKCS5UnPadding(origData)
	// origData = ZeroUnPadding(origData)
	return origData, nil

}

func PKCS5UnPadding(origData []byte) []byte {

	length := len(origData)
	// 去掉最后一个字节 unpadding 次
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]

}
