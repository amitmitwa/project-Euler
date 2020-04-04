package main

import "fmt"

// simplification of (m+n)!/(m!)*(n!) since m=n
func main() {
	var ways int = 1
	for i := 1; i <= 20; i++ {
		ways = (ways * (20 + i)) / i
		//fmt.Println(ways)
	}
	fmt.Println(ways)
}
