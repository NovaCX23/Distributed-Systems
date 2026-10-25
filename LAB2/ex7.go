package main

import "fmt"

func robot(nume string, chTrimite, chPrimeste, observer chan string) {
	chTrimite <- nume + " -> X:10, Y:20"
	chTrimite <- nume + " -> X:15, Y:25"

	close(chTrimite)

	// Citeste automat tot ce vine de la celalalt, pana cand si el da close()
	for pozitie := range chPrimeste {
		fmt.Printf("[%s] a receptionat: %s\n", nume, pozitie)
	}

	observer <- nume + " a terminat sarcina"
}

func main() {
	ch1 := make(chan string, 5)
	ch2 := make(chan string, 5)

	observer := make(chan string)

	// Pornim robotii. R1 scrie in ch1 si citeste din ch2. R2 invers.
	go robot("R1", ch1, ch2, observer)
	go robot("R2", ch2, ch1, observer)

	fmt.Println("Observer:", <-observer)
	fmt.Println("Observer:", <-observer)
}
