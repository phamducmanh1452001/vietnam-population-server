package utils

import (
	"strings"
)

func getSearchKeyArray(key string) []string {
	words := strings.Split(key, " ")

	// length := len(key)
	// for i := 0; i < length-3; i++ {
	// 	words = append(words, strings.Trim(key[i:i+3], " "))
	// }
	return words
}
