package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("ex7.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanWords)

	var words []string
	for scanner.Scan() {
		word := scanner.Text()
		words = append(words, word)
	}
	//fmt.Println(words)
	// De fiecare data , verificam daca citirea s-a finalizat cu succes
	err = scanner.Err()
	if err != nil {
		fmt.Println("Eroare la citire: ", err)
	}

	fr := map[string]int{}

	for _, word := range words {
		fr[word]++
	}

	fmt.Println(fr)
}
