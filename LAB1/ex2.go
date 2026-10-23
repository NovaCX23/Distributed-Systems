package main

import "fmt"

func main() {
	v := [...]int{6, 69, 420, 6, 1, 2, 9, 12, 12, 32, 52, 56, 99}

	ok := true

	for ok {
		ok = false
		for i := range len(v) - 1 {
			if v[i] > v[i+1] {
				ok = true
				v[i], v[i+1] = v[i+1], v[i]
			}
		}
	}

	for i := range v {
		fmt.Printf("%d ", v[i])
	}

}
