package main

import "fmt"

func main() {
	ch := make(chan string)

	go func() {
		ch <- "salut"
	}()
	mesaj := <-ch
	fmt.Println(mesaj)
}
