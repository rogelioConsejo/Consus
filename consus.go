package main

import "math/big"

func main() {
	var f1 = big.NewFloat(.1)
	var f2 = big.NewFloat(.1)
	println(f2.Add(f2, f1).String())
}
