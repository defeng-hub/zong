package main

func lengthOfLastWord(s string) int {
	flag := 0
	length := 0

	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			flag = 1
			continue
		} else {
			if flag == 1 {
				length = 0
				flag = 0
			}
			length++
		}
	}

	return length
}
