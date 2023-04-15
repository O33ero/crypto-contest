package main

import (
	"log"
	"os"
)

var (
	taskFile = ".\\task\\task2"
	maxLen   = 250
)

func main() {
	text, err := getFileContent(taskFile)
	if err != nil {
		return
	}

	guessKeyLenIC(text, maxLen)
}

func guessKeyLenIC(text []byte, maxLen int) {

	for keyLen := 2; keyLen < maxLen; keyLen++ {
		IC := 0.0
		for offset := 0; offset < keyLen; offset++ {
			frequency := countBytesFrequency(text, offset, keyLen)
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

func countBytesFrequency(text []byte, offset int, keyLen int) map[byte]int {
	result := make(map[byte]int)

	for p := offset; p < len(text); p += keyLen {
		c := text[p]
		val, contain := result[c]
		if contain {
			result[c] = val + 1
		} else {
			result[c] = 1
		}
	}
	return result
}

func getFileContent(filePath string) ([]byte, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal("Failed to open file", err.Error())
		return nil, err
	}
	return data, nil
}
