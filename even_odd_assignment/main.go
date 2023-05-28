package main

import "fmt"

func main() {
	nums := makeRange(0, 10)
	for i := range nums {
		parity := calculateParity(i)
		fmt.Println(i, "is", parity)
	}
}

func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}

func calculateParity(num int) string {
	if num%2 == 0 {
		return "even"
	}
	return "odd"
}
