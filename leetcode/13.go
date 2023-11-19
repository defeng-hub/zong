package main

func romanToInt(s string) int {
	var roman = map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	all := 0
	for i := 0; i < len(s)-1; i++ {
		if roman[s[i]] < roman[s[i+1]] {
			all += roman[s[i]] * (-1)
		} else {
			all += roman[s[i]]
		}
	}
	all += roman[s[len(s)-1]]
	return all
}

func max1(num int) (byte, int) {
	if num-1000 >= 0 {
		return 'M', num - 1000
	}
	if num-500 >= 0 {
		return 'D', num - 500
	}
	if num-100 >= 0 {
		return 'C', num - 100
	}
	if num-50 >= 0 {
		return 'L', num - 50
	}
	if num-10 >= 0 {
		return 'X', num - 10
	}
	if num-5 >= 0 {
		return 'V', num - 5
	}
	if num-1 >= 0 {
		return 'I', num - 1
	}
	return '-', 0
}

func intToRoman(num int) string {
	var lnum = num
	var b byte
	var m = make(map[byte]int, 10)
	for {
		b, lnum = max1(lnum)
		_, ok := m[b]
		if ok {
			m[b]++
		} else {
			m[b] = 1
		}
		if lnum == 0 {
			break
		}
	}

	var res string
	if n, ok := m['M']; ok {
		for i := 0; i < n; i++ {
			res = res + "M"
		}
	}

	if n, ok := m['D']; ok {
		for i := 0; i < n; i++ {
			res = res + "D"
		}
	}

	if n, ok := m['C']; ok {
		if n >= 4 {
			res = res + "CD"
		}
		for i := 0; i < n; i++ {
			res = res + "D"
		}
	}
	if n, ok := m['I']; ok {
		if n > 4 {

		}
	}

	return res
}
