package main

import "fmt"

func main() {
	ans := divisible(3, 999) + divisible(5, 999) -divisible(15, 999)
	fmt.Println(ans)

}
func divisible(n, target int) int {
	p := target / n
	return n * (p * (p + 1)) / 2

}
