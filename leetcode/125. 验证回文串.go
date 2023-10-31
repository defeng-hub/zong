package main

import (
	"bytes"
)

func isPalindrome(s string) bool {
	var buffer bytes.Buffer
	for _, v := range s {
		if (v >= 'a' && v <= 'z') || (v >= 'A' && v <= 'Z') {
			buffer.WriteRune(v)
		}
	}

	return false
}
