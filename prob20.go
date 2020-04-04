package main

import (
	"fmt"
	"math/big"
)

func main() {
	var ans big.Int
	ans.MulRange(1, 100)
	s := ans.String()
	sum := 0
	for _, val := range s {
		sum += int(val - '0')
	}
	fmt.Println(&ans, sum)

}
