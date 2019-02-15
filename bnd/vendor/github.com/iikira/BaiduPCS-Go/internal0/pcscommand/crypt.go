package pcscommand

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
)

func AESEncrypt(data []byte) []byte {
	block, err := aes.NewCipher(SK)
	if err != nil {
		return nil
	}
	ecb := cipher.NewCBCEncrypter(block, []byte("RandomInitVector"))
	content := data
	content = PKCS5Padding(content, block.BlockSize())
	crypted := make([]byte, len(content))
	ecb.CryptBlocks(crypted, content)

	return crypted
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func AESDecrypt(crypt []byte) []byte {
	block, err := aes.NewCipher(SK)
	if err != nil {
		return nil
	}
	ecb := cipher.NewCBCDecrypter(block, []byte("RandomInitVector"))
	decrypted := make([]byte, len(crypt))
	ecb.CryptBlocks(decrypted, crypt)

	return PKCS5Trimming(decrypted)
}

func PKCS5Trimming(encrypt []byte) []byte {
	padding := encrypt[len(encrypt)-1]
	return encrypt[:len(encrypt)-int(padding)]
}
