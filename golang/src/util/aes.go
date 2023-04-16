package util

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha1"
	"golang.org/x/crypto/pbkdf2"
	"log"
)

func GetAesCipher(password []byte, salt []byte, keyLength int, iter int) (cipher.Block, error) {
	key := pbkdf2.Key(password, salt, iter, keyLength, sha1.New)
	cipher, err := aes.NewCipher(key)
	if err != nil {
		log.Fatal("Failed to create AES instance", err.Error())
		return nil, err
	}
	return cipher, nil
}
