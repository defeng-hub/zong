package main

import "strings"

func reverseWords(s string) string {
	ss := strings.Trim(s, " ")

	list := strings.Split(ss, " ")

	res := ""
	for i := len(list) - 1; i >= 0; i-- {
		if list[i] == "" || list[i] == " " {
			continue
		}
		res += list[i] + " "
	}
	return res[0 : len(res)-1]
}
