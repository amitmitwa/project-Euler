package main

import "fmt"

func main() {
	for i := 1; i < 500; i++ {
		for j := 1; j < 500; j++ {
			k := 1000 - i - j
			if k*k == j*j+i*i {
				fmt.Println(i * j * k)
				break
			}

		}
	}
}
