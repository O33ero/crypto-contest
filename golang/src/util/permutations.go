package util

func GetNextPermutations(perm []byte, alphabet map[byte]bool) [][]byte {
	for _, c := range perm {
		delete(alphabet, c)
	}

	var result [][]byte
	for k := range alphabet {
		newPerm := append(CreateCopyArray(perm), k)
		result = append(result, newPerm)
	}

	return result
}

func CreateCopyArray(source []byte) []byte {
	var newInstance []byte
	for _, v := range source {
		newInstance = append(newInstance, v)
	}
	return newInstance
}

func CreateCopyMap(source map[byte]bool) map[byte]bool {
	newInstance := make(map[byte]bool)
	for k, v := range source {
		newInstance[k] = v
	}
	return newInstance
}

func GetNextPermutationsWithDuplicates(perm []byte, alphabet []byte) [][]byte {
	for _, c := range perm {
		index := findFirst(c, alphabet)
		alphabet = removeIndex(index, alphabet)
	}

	var result [][]byte
	for _, k := range alphabet {
		newPerm := append(CreateCopyArray(perm), k)
		result = append(result, newPerm)
	}

	return result
}

func removeIndex(index int, array []byte) []byte {
	array = append(array[:index], array[index+1:]...)
	return array
}

func findFirst(value byte, array []byte) int {
	for i := 0; i < len(array); i++ {
		if value == array[i] {
			return i
		}
	}

	return -1
}
