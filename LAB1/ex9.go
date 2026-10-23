package main

import "fmt"

func main() {

	n := 6
	target := 15

	v := []int{2, 7, 11, 8, 4, 9}

	myMap := map[int]int{}

	for i := 0; i < n; i++ {
		_, ok := myMap[target-v[i]]

		if ok {
			fmt.Println(myMap[target-v[i]], i)
		}

		myMap[v[i]] = i
	}

}
