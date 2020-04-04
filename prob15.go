package main

import (
	"fmt"
	"math/big"
)

//the total number of ways = (n+n)!/(n!*n!)

func main() {
	var f40, ans, f20 big.Int
	ans.Div(f40.MulRange(1, 40), f20.MulRange(1, 20))
	//fact.String()
	ways := ans.Div(&ans, f20.MulRange(1, 20))
	fmt.Println(ways)
}
