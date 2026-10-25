package main

import (
	"fmt"
	"math/rand"
	"time"
)

func startRobot(nume string, ch chan string) {
	time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
	ch <- nume
}

func main() {
	robot1 := make(chan string)
	robot2 := make(chan string)

	go startRobot("Robotul A", robot1)
	go startRobot("Robotul B", robot2)

	select {
	case castigator := <-robot1:
		fmt.Println(castigator, "a castigat cursa!")
	case castigator := <-robot2:
		fmt.Println(castigator, "a castigat cursa!")
	}
}
