package main

import (
	"fmt"
	. "golang/src/util"
	"log"
	"os"
	"strings"
)

var (
	outputFile3, _ = os.OpenFile(".//response//response3-Ilya.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
)

func main() {
	var (
		encryptedFilePath = ".\\test\\task3_2"
		basePermutation   = []byte("Wit8Ched29")
	)

	encryptFileStructure, err := GetEncryptedFileStructure(encryptedFilePath)
	if err != nil {
		return
	}
	log.Println(encryptFileStructure.Iter, encryptFileStructure.PlainTextLen, string(encryptFileStructure.Salt))
	BruteAllKeys2(basePermutation, encryptFileStructure)
}

func BruteAllKeys2(basePermutations []byte, fileStruct EncryptFile) {
	permutationSize := len(basePermutations)
	var permutations [][]byte
	var alphabet = make([]byte, len(basePermutations))

	// init permutations
	for i, c := range basePermutations {
		tryDecrypt2(fileStruct, []byte{c})
		permutations = append(permutations, []byte{c})
		alphabet[i] = c
	}

	// start bruteforce
	for len(permutations) > 0 {
		cur := permutations[len(permutations)-1]
		permutations = permutations[:len(permutations)-1]

		nextPermutations := GetNextPermutationsWithDuplicates(cur, CreateCopyArray(alphabet))
		for _, np := range nextPermutations {
			go tryDecrypt2(fileStruct, np)
			if len(np) != permutationSize {
				permutations = append(permutations, np)
			}
		}
	}
}

func tryDecrypt2(fileStruct EncryptFile, cur []byte) {
	decrypt, err := Decrypt(fileStruct, cur)
	if err != nil {
		log.Fatalln("Failed to decrypt")
	}
	s := string(decrypt)
	if strings.Contains(s, "Send") {
		outputString := fmt.Sprintf("%s: %s\n", string(cur), s)
		log.Print(outputString)
		//_, err = outputFile3.WriteString(outputString)
		//if err != nil {
		//	log.Fatalln("Failed to write to file")
		//}
	}
}
