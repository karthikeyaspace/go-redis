package utils

import "crypto/rand"

func GenerateID(n int) string {
	const charset = "abcdefghijklmnopqrstuvwxyz"
	result := make([]byte, n)
	b := make([]byte, n)
	rand.Read(b)
	for i := range b {
		result[i] = charset[int(b[i])%len(charset)]
	}
	return string(result)
}
