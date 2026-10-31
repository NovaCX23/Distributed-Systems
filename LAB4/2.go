package main

import (
	"fmt"
	"sync"
)

// ex 2: flooding pentru grafuri tare conexe
func nodeGraph(id int, in []<-chan int, out []chan<- int, rounds int, wg *sync.WaitGroup, leaderChan chan int) {
	defer wg.Done()

	maxId := id

	for r := 0; r < rounds; r++ {
		// trimitem maximul pe care il stim pe muchiile de iesire
		for _, c := range out {
			c <- maxId
		}

		// citim de pe muchiile de intrare si actualizam daca gasim ceva mai mare
		for _, c := range in {
			v := <-c
			if v > maxId {
				maxId = v
			}
		}
	}

	if maxId == id {
		leaderChan <- id
	}
}

func main() {
	// graf cu 4 noduri -> diam max 4
	rounds := 4

	c01 := make(chan int, 1)
	c12 := make(chan int, 1)
	c23 := make(chan int, 1)
	c30 := make(chan int, 1)
	c13 := make(chan int, 1) // muchie extra pt graf

	var wg sync.WaitGroup
	leaderChan := make(chan int, 4)

	fmt.Println("Pornim algoritmul pe graful tare conex...")

	wg.Add(4)

	// nod, intrari, iesiri
	// 0
	go nodeGraph(10, []<-chan int{c30}, []chan<- int{c01}, rounds, &wg, leaderChan)

	// 1
	go nodeGraph(40, []<-chan int{c01}, []chan<- int{c12, c13}, rounds, &wg, leaderChan)

	// 2
	go nodeGraph(20, []<-chan int{c12}, []chan<- int{c23}, rounds, &wg, leaderChan)

	// 3
	go nodeGraph(30, []<-chan int{c23, c13}, []chan<- int{c30}, rounds, &wg, leaderChan)

	wg.Wait()

	close(leaderChan)

	for lider := range leaderChan {
		fmt.Println("Liderul ales este nodul cu id ul:", lider)
	}
}
