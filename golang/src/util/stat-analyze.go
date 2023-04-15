package util

func CountBytesFrequency(text []byte, offset int, keyLen int) map[byte]int {
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
