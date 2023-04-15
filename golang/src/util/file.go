package util

import (
	"io/fs"
	"log"
	"os"
)

func GetFileContent(filePath string) ([]byte, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal("Failed to open file", err.Error())
		return nil, err
	}
	return data, nil
}

func PutToFile(filePath string, data []byte) {
	err := os.WriteFile(filePath, data, fs.ModePerm)
	if err != nil {
		return
	}
}
