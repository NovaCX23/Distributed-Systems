package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)

	go func() {
		mesaj := <-ch
		fmt.Println("GOroutina a primit:", mesaj)
	}()
	ch <- "salut"
	time.Sleep(100 * time.Millisecond)
}
