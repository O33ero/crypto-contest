package util

func DecodeText(encoded []byte, key []byte) []byte {
	var decoded []byte
	keyLen := len(key)

	for i := 0; i < len(encoded); i++ {
		decoded = append(decoded, encoded[i]^key[i%keyLen])
	}

	return decoded
}
