package main

import "fmt"

func main() {
	sum, i := 0, 1
	number := 0
	for {
		sum += i
		//fmt.Println(sum)
		if prime(sum) >= 500 {
			if prime(sum) > number {
				number = sum
			}
			break
		}
		i++
	}
	fmt.Println(number)
}

func prime(n int) int {
	//	count := 0
	m := make(map[int]int)

	for n%2 == 0 {
		m[2]++
		//fmt.Print(2, " ")
		n = n / 2
	}

	for i := 3; i <= n; i = i + 2 {
		for n%i == 0 {
			m[i]++
			//	fmt.Print(i, " ")
			n = n / i
		}

	}
	total := 1
	//	fmt.Println(m)
	for _, v := range m {
		total *= (v + 1)
	}
	return total

}
