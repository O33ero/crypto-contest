package util

import (
	"encoding/binary"
	"log"
)

type EncryptFile struct {
	Salt         []byte
	Iter         uint32
	PlainTextLen uint32
	EncryptText  []byte
}

func GetEncryptedFileStructure(path string) (EncryptFile, error) {
	content, err := GetFileContent(path)
	if err != nil {
		log.Fatal("Failed to open file")
		return EncryptFile{}, err
	}

	return EncryptFile{
		Salt:         content[0:4],
		Iter:         binary.BigEndian.Uint32(content[4:8]),
		PlainTextLen: binary.BigEndian.Uint32(content[8:12]),
		EncryptText:  content[12:],
	}, nil
}

func Decrypt(encryptStruct EncryptFile, password []byte) ([]byte, error) {
	aesCipher, err := GetAesCipher(password, encryptStruct.Salt, 16, int(encryptStruct.Iter))
	if err != nil {
		log.Fatal("Failed to create AES-128 instance")
		return nil, err
	}

	sourceBytes := encryptStruct.EncryptText
	destBytes := make([]byte, len(sourceBytes))

	size := 16
	for bs, be := 0, size; bs < len(sourceBytes); bs, be = bs+size, be+size {
		aesCipher.Decrypt(destBytes[bs:be], sourceBytes[bs:be])
	}

	return destBytes, nil
}
