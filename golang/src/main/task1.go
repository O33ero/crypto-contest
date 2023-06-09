package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

var (
	basePermutation        = "2A"
	knownPermutationPrefix = "C195B78D364"
	knownSubstring         = ""
	source                 = []rune("а ен жбеко Трр есзтиваял ап дх увооу  ззлооур  впк твкуегч  ойоир ео в се беалуливезо щигг счетон  тойилотд  исрй  кеаолвыо  вн.рС   агнаяи р  р дюеа пфт аар  сс нко и иер пе ьеревдвп ю ытивр геае ынеш сло  лнат  ыцва   е з еобомпзт  уасенч  чзовит.рв а   евют   ко туД с пртдвиеу хява  Вв з ак сньое  НЗ МыОахК. уе п ля жо нжелвто даы же а амслишре  чвт   поас неюомс ыез лноаа.енус  тюювп")
	outFile                = ".//response//response1.txt"
	file, _                = os.OpenFile(outFile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	permuts                = map[rune]int{
		'1': 0,
		'2': 1,
		'3': 2,
		'4': 3,
		'5': 4,
		'6': 5,
		'7': 6,
		'8': 7,
		'9': 8,
		'A': 9,
		'B': 10,
		'C': 11,
		'D': 12,
		'E': 13,
	}

	//а ен жбеко Тр
	//123456789ABCD
)

func main() {
	generatePermutations([]byte(basePermutation))
}

func generatePermutations(basePermutations []byte) {
	permutationSize := len(basePermutations)
	var permutations [][]byte
	alphabet := make(map[byte]bool)

	// init permutations
	for _, c := range basePermutations {
		permutations = append(permutations, []byte{c})
		alphabet[c] = true
	}

	// start bruteforce
	for len(permutations) > 0 {
		cur := permutations[len(permutations)-1]
		permutations = permutations[:len(permutations)-1]

		nextPermutations := getNextPermutations(cur, createCopyMap(alphabet))

		for _, np := range nextPermutations {
			if len(np) == permutationSize {
				brute(np)
			} else {
				permutations = append(permutations, np)
			}
		}
	}
}

func brute(np []byte) bool {
	permutation := np
	permutedText := permuteText(knownPermutationPrefix+string(permutation), source)

	if strings.Contains(permutedText, knownSubstring) {
		sprintf := fmt.Sprintf("%s: %s\n", knownPermutationPrefix+string(permutation), permutedText)
		log.Printf(sprintf)
		_, err := file.WriteString(sprintf)
		if err != nil {
			log.Fatal("Failed to write to file")
			return true
		}
	}
	return false
}

func permuteText(permutation string, source []rune) string {
	pSize := len(permutation)

	var result string
	for i := 0; i < len(source); i += pSize {
		buff := make([]rune, pSize)
		for j, p := range permutation {
			offset, _ := permuts[p]
			buff[j] = source[i+offset]
		}
		result = result + string(buff[:])
	}

	return result
}

func createCopyMap(source map[byte]bool) map[byte]bool {
	newInstance := make(map[byte]bool)
	for k, v := range source {
		newInstance[k] = v
	}
	return newInstance
}

func createCopyArray(source []byte) []byte {
	var newInstance []byte
	for _, v := range source {
		newInstance = append(newInstance, v)
	}
	return newInstance
}

func getNextPermutations(perm []byte, alphabet map[byte]bool) [][]byte {
	for _, c := range perm {
		delete(alphabet, c)
	}

	var result [][]byte
	for k := range alphabet {
		newPerm := append(createCopyArray(perm), k)
		result = append(result, newPerm)
	}

	return result
}
