package main

import (
	"fmt"
	. "golang/src/util"
	"log"
	"os"
)

var (
	outputFile2, _ = os.OpenFile(".//response//response3.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
)

func main() {
	var (
		encryptedFilePath = ".\\task\\task3_test"
		password          = []byte("test")
	)

	encryptFileStructure, err := GetEncryptedFileStructure(encryptedFilePath)
	if err != nil {
		return
	}
	decrypted, err := Decrypt(encryptFileStructure, password)
	if err != nil {
		return
	}

	BruteAllKeysTest(password, encryptFileStructure)
	log.Println(string(decrypted))
}

func BruteAllKeysTest(basePermutations []byte, fileStruct EncryptFile) {
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
			tryDecryptTest(fileStruct, np)
			if len(np) != permutationSize {
				permutations = append(permutations, np)
			}
		}
	}
}

func tryDecryptTest(fileStruct EncryptFile, cur []byte) {
	decrypt, err := Decrypt(fileStruct, cur)
	if err != nil {
		log.Fatalln("Failed to decrypt")
	}
	s := string(decrypt)
	outputString := fmt.Sprintf("%s: %s\n", string(cur), s)
	//log.Print(outputString)
	_, err = outputFile2.WriteString(outputString)
	if err != nil {
		log.Fatalln("Failed to write to file")
	}
}
