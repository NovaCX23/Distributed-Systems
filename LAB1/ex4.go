package main

import "fmt"

func main() {
	v := []float32{18.0, 234.0, 7.0, 0.67, 100.1, 3.14, 0.045, 69.0, 124.0, 0.655}

	var index int
	fmt.Scanf("%d", &index)

	if 0 < index && index < len(v) {
		v = append(v[:index], v[index+1:]...)
	}

	fmt.Println(v)
}
