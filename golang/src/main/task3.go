package main

import (
	"fmt"
	. "golang/src/util"
	"log"
	"os"
)

var (
	outputFile1, _ = os.OpenFile(".//response//response3.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
)

func main() {
	var (
		encryptedFilePath = ".\\task\\task3"
		basePermutation   = []byte("ADnmf380Ew")
	)

	encryptFileStructure, err := GetEncryptedFileStructure(encryptedFilePath)
	if err != nil {
		return
	}
	log.Println(encryptFileStructure.Iter, encryptFileStructure.PlainTextLen, string(encryptFileStructure.Salt))
	BruteAllKeys(basePermutation, encryptFileStructure)
}

func BruteAllKeys(basePermutations []byte, fileStruct EncryptFile) {
	permutationSize := len(basePermutations)
	var permutations [][]byte
	var alphabet = make([]byte, len(basePermutations))

	// init permutations
	for i, c := range basePermutations {
		permutations = append(permutations, []byte{c})
		alphabet[i] = c
	}

	// start bruteforce
	for len(permutations) > 0 {
		cur := permutations[len(permutations)-1]
		permutations = permutations[:len(permutations)-1]

		nextPermutations := GetNextPermutationsWithDuplicates(cur, CreateCopyArray(alphabet))
		for _, np := range nextPermutations {
			tryDecrypt(fileStruct, np)
			if len(np) != permutationSize {
				permutations = append(permutations, np)
			}
		}
	}
}

func tryDecrypt(fileStruct EncryptFile, cur []byte) {
	decrypt, err := Decrypt(fileStruct, cur)
	if err != nil {
		log.Fatalln("Failed to decrypt")
	}
	s := string(decrypt)
	outputString := fmt.Sprintf("%s: %s\n", string(cur), s)
	//log.Print(outputString)
	_, err = outputFile1.WriteString(outputString)
	if err != nil {
		log.Fatalln("Failed to write to file")
	}
}
