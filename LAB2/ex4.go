package main

import "fmt"

func main() {
	ok := make(chan bool)

	go func() {
		fmt.Println("Procesul ruleaza")
		ok <- true
	}()
	<-ok
	fmt.Println("Procesul s-a incheiat")
}
