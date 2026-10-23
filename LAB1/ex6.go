package main

import "fmt"

func main() {

	myMap := map[string]int{
		"Andrei": 190,
		"Maria":  165,
		"Ion":    182,
		"Elena":  172,
		"George": 178,
		"Ana":    165,
	}

	slice := make([]string, 0, len(myMap))

	for nume := range myMap {
		slice = append(slice, nume)
	}

	ok := true
	for ok {
		ok = false
		for i := 0; i < len(myMap)-1; i++ {
			if myMap[slice[i]] < myMap[slice[i+1]] {
				ok = true
				slice[i], slice[i+1] = slice[i+1], slice[i]
			}
		}
	}

	fmt.Println(slice)
}
