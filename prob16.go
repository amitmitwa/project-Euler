package main

import (
	"fmt"
	"math/big"
)

func main() {
	sum := big.NewInt(2)
	//	mult := big.NewInt(2)
	for i := 0; i < 999; i++ {
		sum.Mul(sum, big.NewInt(2))
	}

	fmt.Println(sum)
	s1 := sum.String()
	ans := 0
	for _, val := range s1 {
		ans += int(val - '0')
	}
	fmt.Println(ans)
}
