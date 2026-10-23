// package main

// import "fmt"

// func send(chOut chan string) {
// 	chOut <- "Mesaj de la send"
// }
// func receive(chIn chan string, chOut chan string) {
// 	mesaj := <-chIn
// 	chOut <- mesaj
// }

// func main() {
// 	ch1 := make(chan string)
// 	ch2 := make(chan string)

// 	send(ch1)
// 	receive(ch1, ch2)

// 	fmt.Println(<-ch2)
// }
