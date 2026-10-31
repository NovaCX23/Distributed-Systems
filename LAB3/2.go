package main

import (
	"fmt"
	"math/rand"
	"sort"
	"sync"
)

type LamportClock struct {
	counter int
}

func (l *LamportClock) Tick() {
	l.counter++
}

func (l LamportClock) Send() int {
	return l.counter
}

func (l *LamportClock) Receive(ts int) {
	if ts > l.counter {
		l.counter = ts
	}
	l.counter++
}

type Message struct {
	sender int
	ts     int
}

type Event struct {
	process          int
	internal_counter int
	kind             string // "internal", "send", "recv"
}

func less(a, b Event) bool {
	if a.internal_counter == b.internal_counter {
		// departajare la evenimente concurente
		return a.process < b.process
	}
	return a.internal_counter < b.internal_counter
}

func processNode(id int, in <-chan Message, out chan<- Message, events chan<- Event, wg *sync.WaitGroup) {
	defer wg.Done()
	clock := LamportClock{counter: 0}

	// tick random
	randomTicks := rand.Intn(3) + 1
	for i := 0; i < randomTicks; i++ {
		clock.Tick()
		events <- Event{process: id, internal_counter: clock.counter, kind: "internal"}
	}

	// send
	clock.Tick()
	ts := clock.Send()
	events <- Event{process: id, internal_counter: ts, kind: "send"}
	out <- Message{sender: id, ts: ts}

	// recv
	msg := <-in
	clock.Receive(msg.ts)
	events <- Event{process: id, internal_counter: clock.counter, kind: "recv"}
}

func main() {
	// topologie inel
	c12 := make(chan Message, 1)
	c23 := make(chan Message, 1)
	c31 := make(chan Message, 1)

	eventsChan := make(chan Event, 20)
	var wg sync.WaitGroup

	wg.Add(3)

	// 1->2->3->1
	go processNode(1, c31, c12, eventsChan, &wg)
	go processNode(2, c12, c23, eventsChan, &wg)
	go processNode(3, c23, c31, eventsChan, &wg)

	wg.Wait()
	close(eventsChan)

	var allEvents []Event
	for e := range eventsChan {
		allEvents = append(allEvents, e)
	}

	// sortare
	sort.Slice(allEvents, func(i, j int) bool {
		return less(allEvents[i], allEvents[j])
	})

	for _, e := range allEvents {
		fmt.Printf("P%d | Timp: %d | Tip: %s\n", e.process, e.internal_counter, e.kind)
	}
}
