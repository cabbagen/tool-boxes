package utils

import (
	"log"
	"fmt"
	"bytes"
	"crypto/aes"
	"crypto/md5"
	"encoding/hex"
	"encoding/base64"
)

// base64
func Base64EncodeFunc (str string) string {
	encoded := base64.StdEncoding.EncodeToString([]byte(str))
	
	return encoded
}

func Base64DecodeFunc (str string) string {
	decoded, error := base64.StdEncoding.DecodeString(str)
	
	if error != nil {
		fmt.Println("decode error: ", error)
		return ""
	}
	
	return string(decoded)
}

// 十六进制
func HexEncodeFunc (str string) string {
	encodestring := hex.EncodeToString([]byte(str))
	
	return encodestring
}

func HexDecodeFunc (str string) string {
	decode, error := hex.DecodeString(str)
	
	if error != nil {
		fmt.Println("decode error: ", error)
		return ""
	}
	
	return string(decode)
}

// AES ECB
func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext) % blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	
	return append(ciphertext, padtext...)
}

func PKCS5UnPadding(originData []byte) []byte {
	length := len(originData)
	unpadding := int(originData[length - 1])
	
	return originData[:(length - unpadding)]
}

func AESECBEncode(key []byte, plainText string) string {
	
	block, error := aes.NewCipher(key)
	
	if error != nil {
		log.Fatal(error)
	}
	
	source := PKCS5Padding([]byte(plainText), block.BlockSize())
	
	dist := make([]byte, len(source))
	
	for i := 0; i < len(source); i += block.BlockSize() {		
		block.Encrypt(dist[i: i + block.BlockSize()], source[i: i + block.BlockSize()])
	}
	
	return hex.EncodeToString(dist)
}

func AESECBDecode(key []byte, secretText string) string {
	block, error := aes.NewCipher(key)
	
	if error != nil {
		log.Fatal(error)
	}
	
	source, error := hex.DecodeString(secretText)
	
	if error != nil {
		log.Fatal(error)
	}
	
	dist := make([]byte, len(source))
	
	for i := 0; i < len(source); i += block.BlockSize() {
		block.Decrypt(dist[i: i + block.BlockSize()], source[i: i + block.BlockSize()])
	}
	
	return string(PKCS5UnPadding(dist))
}

// MD5 签名
func MD5(str string) string {

	var sign bytes.Buffer
	
	if _, error := fmt.Fprintf(&sign, "%x", md5.Sum([]byte(str))); error != nil {
		log.Fatal(error)
	}
	
	return sign.String()
}













