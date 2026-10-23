package main

import "fmt"

func main() {

	angajati := map[string]int{
		"Ana Popescu":      1,
		"Mihai Ionescu":    2,
		"Bogdan Radu":      1,
		"Elena Dumitrescu": 3,
		"Cristian Vasile":  2,
		"Ioana Enache":     1,
	}

	departament := map[int][]string{}

	for key, val := range angajati {
		departament[val] = append(departament[val], key)
	}

	fmt.Println(departament)

}
