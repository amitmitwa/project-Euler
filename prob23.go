package main

import "fmt"

func main() {
	var arr []int
	check := make([]bool, 28124)
	var finalAns int
	for i := 1; i <= 28123; i++ {
		if d(i) > i {
			arr = append(arr, i)
			//finalAns += i
		}

	}
	fmt.Println(arr, len(arr))
	for i := 0; i < len(arr); i++ {
		for j := i; j < len(arr); j++ {
			if arr[i]+arr[j] <= 28123 {
				check[arr[i]+arr[j]] = true
			}
		}
	}
	for i := 1; i <= 28123; i++ {
		if !check[i] {
			finalAns += i
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
