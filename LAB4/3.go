package main

import (
	"fmt"
	"sync"
)

// ex 3: flooding de medie pentru vot majoritar
func nodeMajority(initialVal float64, in []<-chan float64, out []chan<- float64, rounds int, wg *sync.WaitGroup, resultChan chan float64) {
	defer wg.Done()

	x := initialVal
	d := float64(len(in)) // grad intern

	for r := 0; r < rounds; r++ {
		// trimitem val curenta
		for _, c := range out {
			c <- x
		}

		// primim val de la vecini si le adunam
		sum := x
		for _, c := range in {
			sum += <-c
		}

		// actualizam starea locala med arit
		x = sum / (d + 1.0)
	}

	resultChan <- x
}

func main() {
	rounds := 20

	// ca sa facem o matrice dublu stocastica folosim un inel bidirectional
	c01 := make(chan float64, 1)
	c10 := make(chan float64, 1)

	c12 := make(chan float64, 1)
	c21 := make(chan float64, 1)

	c23 := make(chan float64, 1)
	c32 := make(chan float64, 1)

	c30 := make(chan float64, 1)
	c03 := make(chan float64, 1)

	var wg sync.WaitGroup
	resultChan := make(chan float64, 4)

	fmt.Println("Pornim votul...")

	wg.Add(4)

	// nodurile voteaza 1 sau 0
	go nodeMajority(1.0, []<-chan float64{c10, c30}, []chan<- float64{c01, c03}, rounds, &wg, resultChan)
	go nodeMajority(1.0, []<-chan float64{c01, c21}, []chan<- float64{c10, c12}, rounds, &wg, resultChan)
	go nodeMajority(0.0, []<-chan float64{c12, c32}, []chan<- float64{c21, c23}, rounds, &wg, resultChan)
	go nodeMajority(1.0, []<-chan float64{c23, c03}, []chan<- float64{c32, c30}, rounds, &wg, resultChan)

	wg.Wait()
	close(resultChan)

	fmt.Println("Valorile de consens la final:")
	for val := range resultChan {
		fmt.Printf("%.3f\n", val)

		if val > 0.5 {
			fmt.Println(" -> Rezultatul pentru acest nod: MAJORITATE PENTRU DA")
		} else {
			fmt.Println(" -> Rezultatul pentru acest nod: MAJORITATE PENTRU NU")
		}
	}
}
