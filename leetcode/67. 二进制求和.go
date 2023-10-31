package main

import "math/big"

func addBinary(a string, b string) string {
	ai, _ := new(big.Int).SetString(a, 2)
	bo, _ := new(big.Int).SetString(b, 2)
	res := ai.Add(ai, bo)
	return res.Text(10)
}
