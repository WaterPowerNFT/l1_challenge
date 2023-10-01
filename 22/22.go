package main

import (
	"fmt"
	"math/big"
)

func SumBI(a *big.Int, b *big.Int) big.Int {
	var c big.Int
	c.Add(a, b)
	return c
}

func SubBI(a *big.Int, b *big.Int) big.Int {
	var c big.Int
	c.Sub(a, b)
	return c
}

func MulBI(a *big.Int, b *big.Int) big.Int {
	var c big.Int
	c.Mul(a, b)
	return c
}

func DivBI(a *big.Int, b *big.Int) (big.Int, big.Int) {
	var c, mod big.Int
	c.Div(a, b)
	mod.Mod(a, b)
	return c, mod
}

func main() {

	var x big.Int
	x.SetString("999999999999999999999999999999999999999999999999999999999999999999999999999", 10)
	var y big.Int
	y.SetString("888888888888888888888888888888888888888888888888888888888888888888888888888", 10)
	bi_sum := SumBI(&x, &y)
	fmt.Println("sum is", bi_sum.String())
	bi_sub := SubBI(&x, &y)
	fmt.Println("sub is", bi_sub.String())
	bi_mul := MulBI(&x, &y)
	fmt.Println("mul is", bi_mul.String())
	bi_div, bi_div_mod := DivBI(&x, &y)
	fmt.Println("div and mod is", bi_div.String(), bi_div_mod.String())
}
