package main

import (
	. "golang/src/util"
	"log"
)

var (
	maxLen = 250
)

func main() {
	text, err := GetFileContent(".\\task\\task2")
	if err != nil {
		return
	}

	guessKeyLenIC(text, maxLen)
}

func guessKeyLenIC(text []byte, maxLen int) {

	for keyLen := 2; keyLen < maxLen; keyLen++ {
		IC := 0.0
		for offset := 0; offset < keyLen; offset++ {
			frequency := CountBytesFrequency(text, offset, keyLen)
			textLen := float64(len(text))
			generalFrequency := float64(calcGeneralFrequency(frequency))
			offsetIC := generalFrequency / (textLen * (textLen - 1))

			IC += offsetIC
		}
		log.Printf("Keylen: %2d, IC: %.9f", keyLen, IC)
	}
}

func calcGeneralFrequency(frequency map[byte]int) int {
	sum := 0
	for _, v := range frequency {
		sum += v * (v - 1)
	}
	return sum
}
