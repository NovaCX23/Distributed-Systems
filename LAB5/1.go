package main

import "fmt"

const MAX_ITER = 1000
const R = 16.0
const E = 1e-6

type Punct struct {
	X, Y float64
}

func dist(a, b Punct) float64 {
	return ((a.X-b.X)*(a.X-b.X) + (a.Y-b.Y)*(a.Y-b.Y))
}

func main() {
	pos := []Punct{
		{0, 0},
		{10, 0},
		{20, 5},
		{20, 20},
		{5, 15},
		{0, 10},
	}
	iter := 0
	for len(pos) > 1 && iter < MAX_ITER {
		nou := make([]Punct, len(pos))

		for i := range pos {
			sx, sy := pos[i].X, pos[i].Y
			cnt := 1
			for j := range pos {
				if i != j && dist(pos[i], pos[j]) < R*R {
					sx += pos[j].X
					sy += pos[j].Y
					cnt++
				}
			}
			nou[i] = Punct{sx / float64(cnt), sy / float64(cnt)}
		}
		var rez []Punct
		for _, p := range nou {
			gasit := false
			for _, q := range rez {
				if dist(p, q) < E {
					gasit = true
					break
				}
			}
			if !gasit {
				rez = append(rez, p)
			}
		}
		pos = rez
		iter++
	}
	fmt.Println("Iteratii:", iter)
	if len(pos) == 1 {
		fmt.Printf("Rendezvous in punctul(%.4f, %.4f)\n:", pos[0].X, pos[0].Y)
	} else {
		fmt.Println("Nu avem rendezvous")
		for _, p := range pos {
			fmt.Printf("Cluster in punctul(%.4f, %.4f)\n:", p.X, p.Y)
		}
	}
}
