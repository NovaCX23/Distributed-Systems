package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"strings"
)

func diametru(adj [][]int) int {
	n := len(adj)
	diam := 0

	for src := 0; src < n; src++ {
		// BFS din nodul src
		dist := make([]int, n)
		for i := range dist {
			dist[i] = -1
		}

		dist[src] = 0
		queue := []int{src}

		for len(queue) > 0 {
			u := queue[0]
			queue = queue[1:]
			for _, v := range adj[u] {
				if dist[v] == -1 {
					dist[v] = dist[u] + 1
					queue = append(queue, v)
				}
			}
		}

		for _, d := range dist {
			if d == -1 {
				return -1 // avem un nod la care nu putem ajunge din src
			}
			if d > diam {
				diam = d
			}
		}
	}
	return diam
}

// Info -> pereche (id, valoare) retinuta in memoria unui nod
type Info struct {
	ID    int
	Value int
}

// Memory -> multime token uri (reuniune prin inserare)
type Memory map[Info]struct{}

type Node struct {
	ID        int
	Neighbors []int
	M         Memory
}

// citeste o linie si o transforma in lista de int uri
func citesteLinie(sc *bufio.Scanner) []int {
	var rez []int
	if !sc.Scan() {
		return rez
	}
	for _, x := range strings.Fields(sc.Text()) {
		v, _ := strconv.Atoi(x)
		rez = append(rez, v)
	}
	return rez
}

func main() {
	sc := bufio.NewScanner(os.Stdin)

	prima := citesteLinie(sc)
	n, s := prima[0], prima[1]

	tokens := citesteLinie(sc)

	// listele de adiacenta
	nodes := make([]Node, n)
	adj := make([][]int, n)
	for i := 0; i < n; i++ {
		adj[i] = citesteLinie(sc)
		nodes[i] = Node{ID: i, Neighbors: adj[i], M: Memory{}}
		// M_i(0) = {(i, v_i)}
		nodes[i].M[Info{i, tokens[i]}] = struct{}{}
	}

	diam := diametru(adj)
	if diam == -1 {
		fmt.Println("Graful nu este tare conex")
		return
	}
	fmt.Println("Diametrul grafului:", diam)

	crashed := make([]bool, n)
	defecte := 0
	t := 0

	for {
		// decidem random daca un nod intra in crash la iteratia asta
		// (pastram mereu macar un nod corect)
		if defecte < s && defecte < n-1 && rand.Intn(2) == 0 {
			x := rand.Intn(n)
			for crashed[x] {
				x = rand.Intn(n)
			}
			crashed[x] = true
			defecte++
			fmt.Printf("Iteratia %d: nodul %d a intrat in crash\n", t, x)
		}

		// fiecare nod corect trimite M_i(t) la vecini
		// U[j] = reuniunea memoriilor primite de nodul j
		U := make([]Memory, n)
		for i := range U {
			U[i] = Memory{}
		}
		for i := 0; i < n; i++ {
			if crashed[i] {
				continue // nodul in crash nu mai trimite nimic
			}
			for _, j := range nodes[i].Neighbors {
				for info := range nodes[i].M {
					U[j][info] = struct{}{}
				}
			}
		}

		// M_i(t+1) = M_i(t) U U
		for i := 0; i < n; i++ {
			if crashed[i] {
				continue
			}
			for info := range U[i] {
				nodes[i].M[info] = struct{}{}
			}
		}

		if t > (s+1)*diam {
			break
		}

		t++
	}

	fmt.Println("Algoritmul s-a oprit la iteratia:", t)

	// afisam memoria primului nod corect, sortata dupa id
	for i := 0; i < n; i++ {
		if !crashed[i] {
			var lista []Info
			for info := range nodes[i].M {
				lista = append(lista, info)
			}
			sort.Slice(lista, func(a, b int) bool {
				return lista[a].ID < lista[b].ID
			})

			fmt.Printf("Memoria nodului %d:\n", i)
			for _, info := range lista {
				fmt.Printf("(%d, %d) ", info.ID, info.Value)
			}
			fmt.Println()
			break
		}
	}
}
