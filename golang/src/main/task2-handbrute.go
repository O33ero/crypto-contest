package main

import (
	"encoding/hex"
	. "golang/src/util"
	"log"
)

var (
	key = "e697a004edbac0ca37c398aa25d091daee2f7ec2b0ca91e5"
	//     T h e | | | a l g o r y a u k s | | | f i r s t
	//     | | | a u t h e n t i s t i c | | | a | | | m e
	//     s s a g e | | t o | | | k c r i f y | | | t h e
	//     o r i g i n . | | | U s e | | |
)

func main() {
	keyHex, _ := hex.DecodeString(key)

	text, err := GetFileContent(".\\task\\task2_3")
	if err != nil {
		return
	}

	decodedHex := DecodeText(text, keyHex)
	log.Printf("Key: %s\n", keyHex)
	log.Printf("Decoded text (Hex): %x\n", decodedHex)
	PutToFile(".\\response\\response2", decodedHex)
}
