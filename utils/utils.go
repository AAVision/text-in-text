package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"

	"golang.org/x/crypto/pbkdf2"
)

func AESEncrypt(coverText []byte) []byte {

	key := []byte("passphrasewhichneedstobe32bytes!")

	c, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println(err)
	}

	gcm, err := cipher.NewGCM(c)

	if err != nil {
		fmt.Println(err)
	}

	nonce := make([]byte, gcm.NonceSize())

	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		fmt.Println(err)
	}

	return (gcm.Seal(nonce, nonce, coverText, nil))
}

func AESDecrypt(ciphertext []byte) string {
	key := []byte("passphrasewhichneedstobe32bytes!")
	c, err := aes.NewCipher(key)

	if err != nil {
		fmt.Println(err)
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		fmt.Println(err)
	}

	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		fmt.Println(err)
	}

	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		fmt.Println(err)
	}

	return string(plaintext)

}
func Encrypt(plainText, password string) (string, error) {
	salt := make([]byte, 8)
	if _, err := io.ReadFull(rand.Reader, salt); err != nil {
		return "", err
	}

	key := pbkdf2.Key([]byte(password), salt, 10000, 32, sha256.New)

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	iv := make([]byte, aes.BlockSize)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	mode := cipher.NewCBCEncrypter(block, iv)

	plainTextBytes := pkcs7Pad([]byte(plainText), aes.BlockSize)
	cipherText := make([]byte, len(plainTextBytes))
	mode.CryptBlocks(cipherText, plainTextBytes)

	finalCipherText := append(salt, iv...)
	finalCipherText = append(finalCipherText, cipherText...)

	return base64.StdEncoding.EncodeToString(finalCipherText), nil
}

// Decrypt decrypts ciphertext using AES with a password.
func Decrypt(cipherTextBase64, password string) (string, error) {
	cipherText, err := base64.StdEncoding.DecodeString(cipherTextBase64)
	if err != nil {
		return "", err
	}

	salt := cipherText[:8]
	iv := cipherText[8 : 8+aes.BlockSize]
	cipherText = cipherText[8+aes.BlockSize:]

	key := pbkdf2.Key([]byte(password), salt, 10000, 32, sha256.New)

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	mode := cipher.NewCBCDecrypter(block, iv)

	plainText := make([]byte, len(cipherText))
	mode.CryptBlocks(plainText, cipherText)

	plainText, err = pkcs7Unpad(plainText)
	if err != nil {
		return "", err
	}

	return string(plainText), nil
}

// pkcs7Pad pads the data to be encrypted to a multiple of blockSize
func pkcs7Pad(data []byte, blockSize int) []byte {
	padding := blockSize - len(data)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padText...)
}

// pkcs7Unpad removes padding from the decrypted data
func pkcs7Unpad(data []byte) ([]byte, error) {
	length := len(data)
	if length == 0 {
		return nil, fmt.Errorf("pkcs7: invalid padding size")
	}

	unpadding := int(data[length-1])
	if unpadding > length || unpadding > aes.BlockSize {
		return nil, fmt.Errorf("pkcs7: invalid padding")
	}

	return data[:length-unpadding], nil
}
