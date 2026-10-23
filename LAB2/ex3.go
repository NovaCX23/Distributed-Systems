package main

import "fmt"

func main() {
	ch := make(chan string, 2)

	ch <- "GO"
	ch <- "LANG"

	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
