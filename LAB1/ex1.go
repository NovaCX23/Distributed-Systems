package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)

	switch {
	case x < 2:
		fmt.Println("Nu e prim")
		return
	case x == 2:
		fmt.Println("E prim")
		return
	case x%2 == 0:
		fmt.Println("Nu e prim")
		return
	}

	for d := 3; d*d <= x; d += 2 {
		if x%d == 0 {
			fmt.Println("Nu e prim")
			return
		}
	}

	fmt.Println("E prim")
}
