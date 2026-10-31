package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Clock struct {
	offset time.Duration
}

func (c Clock) Now() time.Time {
	return time.Now().Add(c.offset)
}

func estimate(t1, t2, t3, t4 time.Time) (theta, delta time.Duration) {
	// offset: theta = ((T2 - T1) + (T3 - T4)) / 2
	theta = (t2.Sub(t1) + t3.Sub(t4)) / 2

	// round-trip: delta = (T4 - T1) - (T3 - T2)
	delta = (t4.Sub(t1)) - (t3.Sub(t2))

	return theta, delta
}

func measure(c_clock, s_clock Clock) (theta, delta time.Duration) {
	// Request
	t1 := c_clock.Now()
	// intarziere retea ->
	time.Sleep(3*time.Millisecond + time.Duration(rand.Intn(10))*time.Millisecond)

	t2 := s_clock.Now()
	// simulam procesarea pe server
	time.Sleep(5*time.Millisecond + time.Duration(rand.Intn(20))*time.Millisecond)
	t3 := s_clock.Now()

	// Response
	// intarziere retea <-
	time.Sleep(3*time.Millisecond + time.Duration(rand.Intn(10))*time.Millisecond)
	t4 := c_clock.Now()

	return estimate(t1, t2, t3, t4)
}

func main() {
	serverClock := Clock{offset: 0}
	initialOffset := 500 * time.Millisecond
	clientClock := Clock{offset: initialOffset}

	fmt.Printf("Offset: %v\n\n", clientClock.offset)

	fmt.Println("a) O singura masuratoare")
	thetaA, deltaA := measure(clientClock, serverClock)

	// corectie
	clientClockA := Clock{offset: initialOffset + thetaA}

	fmt.Printf("Delta: %v\n", deltaA)
	fmt.Printf("Theta: %v\n", thetaA)
	fmt.Printf("Offset + corectie: %v\n\n", clientClockA.offset)

	fmt.Println("b) 10 masuratori")

	minDelta := time.Hour
	var bestTheta time.Duration

	for i := 0; i < 10; i++ {
		theta, delta := measure(clientClock, serverClock)

		if delta < minDelta {
			minDelta = delta
			bestTheta = theta
		}
	}

	clientClockB := Clock{offset: initialOffset + bestTheta}

	fmt.Printf("Delta min: %v\n", minDelta)
	fmt.Printf("Theta: %v\n", bestTheta)
	fmt.Printf("Offset + corectie: %v\n", clientClockB.offset)
}
