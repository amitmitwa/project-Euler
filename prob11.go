package main

import (
	"fmt"
)

func main() {
	var a [20][20]int

	for i := 0; i < 20; i++ {
		for j := 0; j < 20; j++ {
			fmt.Scan(&a[i][j])
		}
	}
	x := row(a)
	y := column(a)
	z := diagonal1(a)
	w := diagonal2(a)

	if x > y && x > z && x > w {
		fmt.Println(x)
	}
	if y > x && y > z && y > w {
		fmt.Println(y)
	}
	if z > y && z > x && z > w {
		fmt.Println(z)
	}
	if w > x && w > y && w > z {
		fmt.Println(w)
	}

}

func row(a [20][20]int) int {
	sum, maxSum := 1, 0
	for i := 0; i < 20; i++ {
		for j := 0; j < 16; j++ {
			sum = a[i][j] * a[i][j+1] * a[i][j+2] * a[i][j+3]
			if sum > maxSum {
				maxSum = sum
			}
		}

	}
	return maxSum
}

func column(a [20][20]int) int {
	sum, maxSum := 1, 0
	for i := 0; i < 16; i++ {
		for j := 0; j < 20; j++ {
			sum = a[i][j] * a[i+1][j] * a[i+2][j] * a[i+3][j]
			if sum > maxSum {
				maxSum = sum

			}
		}

	}
	return maxSum
}

func diagonal1(a [20][20]int) int {
	sum, maxSum := 1, 0
	for i := 0; i < 16; i++ {
		for j := 0; j < 16; j++ {
			sum = a[i][j] * a[i+1][j+1] * a[i+2][j+2] * a[i+3][j+3]
			if sum > maxSum {
				maxSum = sum

			}
		}

	}
	return maxSum
}
func diagonal2(a [20][20]int) int {
	sum, maxSum := 1, 0
	for i := 0; i < 16; i++ {
		for j := 3; j < 20; j++ {
			sum = a[i][j] * a[i+1][j-1] * a[i+2][j-2] * a[i+3][j-3]
			if sum > maxSum {
				maxSum = sum
			}
		}

	}
	return maxSum
}
