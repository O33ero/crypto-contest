package main

import (
	"encoding/hex"
	. "golang/src/util"
	"log"
)

var (
	key = "85fd5d8c671417e9eb413c719b67c4bd53c6d60c7dfcf9d3feaf15ac99a7c0"
)

func main() {
	keyHex, _ := hex.DecodeString(key)

	text, err := GetFileContent(".\\task\\task2")
	if err != nil {
		return
	}

	decodedHex := DecodeText(text, keyHex)
	log.Printf("Key: %s\n", keyHex)
	log.Printf("Decoded text (Hex): %x\n", decodedHex)
	PutToFile(".\\response\\response2", decodedHex)
}
