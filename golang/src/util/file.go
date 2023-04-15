package util

import (
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
