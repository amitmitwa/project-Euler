package main

import "fmt"

func main() {
	max, num := 0, 0
	for i := 999999; i > 500000; i = i - 2 {
		//fmt.Println(check(i))
		if check(i) > max {
			max, num = check(i), i
		}

	}
	fmt.Println(max, num)

}
func check(x int) int {
	count := 0
	for x != 1 {
		count++
		if x%2 == 1 {
			x = 3*x + 1

		} else {
			x = x / 2

		}
	}
	return count + 1
}
