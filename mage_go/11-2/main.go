package main

import "fmt"

func main() {
	//s := "abcdwdf,wds"
	//
	//fmt.Printf("%#v\n", strings.Split(s, ","))
	//
	//fmt.Printf("%d\n", strings.Index(s, "wdf"))
	//fmt.Println(strings.Join(strings.Split(s, ","), "----"))

	s := "12abc王德峰aaa"
	arr := []byte(s)
	//brr := []rune(s)
	for _, ele := range arr {
		fmt.Printf("%d-", ele)
	}
	fmt.Printf("\narr len %d, str len %d\n", len(arr), len(s))
	fmt.Printf("s[1]=%d, %c", s[1], s[7])
}
