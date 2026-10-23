package main

import (
	"fmt"
)

func main() {
	slice := []string{"Laboratorul", "de", "sisteme", "distribuite"}
	n := len(slice)

	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}

	fmt.Println("Slice inversat:", slice)
}
