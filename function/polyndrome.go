package function

import (
	"strings"
)

func Polyndrome(word1 string) bool {

	var trimWord1 = strings.ReplaceAll(word1, " ", "")

	// to lowercase 
	trimWord1 = strings.ToLower(trimWord1)
	var reverseOfTrimWord1 string  = ""
	for i := len(trimWord1)-1 ; i >= 0 ; i-- {
		reverseOfTrimWord1 += string(trimWord1[i])
	}

	if reverseOfTrimWord1 == trimWord1 {
		return true
	}
	return false
}

