package main

import "fmt"

func main() {
	finalAns := 0
	for i := 2; i < 10000; i++ {
		if i == d(d(i)) {
			fmt.Println(i, d(i), d(d(i)))
			finalAns += i
		}
		if i == d(i) {
			finalAns -= i
		}

	}

	fmt.Println(finalAns)
}
func d(n int) int {
	temp := n
	mapp := make(map[int]int)
	for n%2 == 0 {
		n = n / 2
		mapp[2]++

	}
	//	fmt.Println(n)
	for i := 3; i <= n; i = i + 2 {
		for n%i == 0 {
			n = n / i
			mapp[i]++

		}
	}
	//fmt.Println(mapp)
	sum, mul, sumOfD := 1, 1, 1
	for k, v := range mapp {
		for i := 0; i < v; i++ {
			mul *= k
			sum += mul

		}
		//fmt.Println(sum, k, v)
		sumOfD *= sum
		sum, mul = 1, 1

	}
	return sumOfD - temp
}
