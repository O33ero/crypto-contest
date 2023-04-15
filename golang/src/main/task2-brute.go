package main

import (
	"encoding/hex"
	. "golang/src/util"
	"log"
	"math"
)

var (
	keyLen = 31
)

func main() {

	text, err := GetFileContent(".\\task\\task2")
	if err != nil {
		return
	}

	keys := generateKeys(text, keyLen)
	for _, k := range keys {
		keyString := hex.EncodeToString(k)
		decodedHex := DecodeText(text, k)

		log.Printf("Possible key(string): %s\n", keyString)
		log.Printf("Decoded text (hex): %s\n", hex.EncodeToString(decodedHex))
	}
}

func generateKeys(text []byte, keyLen int) [][]byte {
	var (
		allChars = initAllBytes()
		keys     [][]byte
	)
	// init dataset

	for _, c := range allChars {
		gKeys := guessKeys(text, c, keyLen)
		for _, k := range gKeys {
			keys = append(keys, k)
		}
	}

	return keys
}

func initAllBytes() [256]byte {
	var allBytes [256]byte
	for i := 0; i < 256; i++ {
		allBytes[i] = byte(i)
	}
	return allBytes
}

func guessKeys(text []byte, mostChar byte, keyLen int) [][]byte {
	keysPossibleBytes := make([][]byte, keyLen)

	for offset := 0; offset < keyLen; offset++ {
		frequency := CountBytesFrequency(text, offset, keyLen)
		maxFreq := getMaxValueInMap(frequency)
		for f := range frequency {
			if float64(frequency[f]) >= maxFreq {
				keysPossibleBytes[offset] = append(keysPossibleBytes[offset], f^mostChar)
			}
		}
	}

	keys := getNextCombination(keysPossibleBytes, []byte{}, 0)
	return keys
}

func getMaxValueInMap(frequency map[byte]int) float64 {
	maxFreq := 0.0
	for _, f := range frequency {
		maxFreq = math.Max(maxFreq, float64(f))
	}
	return maxFreq
}

func getNextCombination(alphabet [][]byte, keyPart []byte, offset int) [][]byte {
	var keys [][]byte
	if offset >= len(alphabet) {
		return [][]byte{keyPart}
	}
	for _, c := range alphabet[offset] {
		for _, k := range getNextCombination(alphabet, append(keyPart, c), offset+1) {
			keys = append(keys, k)
		}
	}
	return keys
}
