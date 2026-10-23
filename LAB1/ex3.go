package main

import "fmt"

func main() {
	v := [...]int{1, 2, 3, 4, 5, 6, 7, 8}

	var even []int

	for _, val := range v {
		if val%2 == 0 {
			even = append(even, val)
		}
	}

	fmt.Println("Nr:", even)
}
