package main

import (
	"fmt"
	"math/big"
)

func main() {
	var a, b big.Int

	a.Exp(big.NewInt(2), big.NewInt(30), nil)
	b.Exp(big.NewInt(3), big.NewInt(30), nil)

	fmt.Println(plus(&a, &b))
	fmt.Println(minus(&a, &b))
	fmt.Println(mul(&a, &b))
	fmt.Println(div(&a, &b))

}

func plus(a, b *big.Int) *big.Int {
	return new(big.Int).Add(a, b)
}

func minus(a, b *big.Int) *big.Int {
	return new(big.Int).Sub(a, b)
}

func mul(a, b *big.Int) *big.Int {
	return new(big.Int).Mul(a, b)
}

func div(a, b *big.Int) *big.Int {
	return new(big.Int).Div(a, b)
}
