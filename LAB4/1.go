package main

import (
	"fmt"
	"sync"
)

type Msg struct {
	Value int
	Done  bool // true cand s a gasit liderul
}

func node(id int, in <-chan Msg, out chan Msg, isLeader *bool, wg *sync.WaitGroup, leaderChan chan int) {
	defer wg.Done()

	// trimitem id ul propriu ca sa pornim algoritmul
	out <- Msg{Value: id, Done: false}

	for msg := range in {
		// daca s a ales liderul, dam mesajul mai departe si iesim
		if msg.Done {
			if !*isLeader {
				out <- msg
			}
			break
		}

		// dam mai departe doar id urile mai mari
		if msg.Value > id {
			out <- msg
		} else if msg.Value == id {
			// id ul meu a facut un tur complet, eu sunt liderul
			*isLeader = true
			leaderChan <- id

			// trimit semnalul de oprire in inel
			out <- Msg{Value: id, Done: true}
		}
	}
}

func RingFlooding(ids []int) int {
	n := len(ids)
	if n == 0 {
		return -1
	}

	// canale pentru legaturile dintre noduri
	chans := make([]chan Msg, n)
	for i := range chans {
		chans[i] = make(chan Msg, n)
	}

	var wg sync.WaitGroup
	isLeaders := make([]bool, n)
	leaderChan := make(chan int, 1)

	for i := 0; i < n; i++ {
		wg.Add(1)
		go node(ids[i], chans[i], chans[(i+1)%n], &isLeaders[i], &wg, leaderChan)
	}

	leaderId := <-leaderChan

	wg.Wait()

	return leaderId
}

func main() {
	// hardcodam niste id uri pentru test
	ids := []int{15, 3, 27, 8, 12, 5}

	fmt.Println("Pornim alegerea liderului in inel cu id urile:", ids)

	lider := RingFlooding(ids)

	fmt.Println("Liderul ales este nodul cu id ul:", lider)
}
