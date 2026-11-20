package main

import (
	"fmt"
	"math/rand"
)

// Message reprezinta un mesaj transmis intre noduri in Runda 1 sau 2
type Message struct {
	SenderID int
	Round    int // 1 pentru starea curenta, 2 pentru valoarea Tie
	Value    int // 0 sau 1
}

type Node struct {
	id          int
	x           int // starea curenta x_i(t)
	s           int
	n           int  // numarul de noduri din sistem
	isByzantine bool // comportamentul la iteratia curenta
	inbox       chan Message
}

func Bcast(nodes []Node, from int, round int, val int) {
	for j := range nodes {
		nodes[j].inbox <- Message{SenderID: from, Round: round, Value: val}
	}
}

// runda 1, nod bizantin -> val aleatorie
func MaliciousBcastState(nodes []Node, from int) {
	for j := range nodes {
		nodes[j].inbox <- Message{SenderID: from, Round: 1, Value: rand.Intn(2)}
	}
}

// runda 2, nod bizantin -> val Tie contradicitorii
func MaliciousBcastTie(nodes []Node, from int) {
	for j := range nodes {
		nodes[j].inbox <- Message{SenderID: from, Round: 2, Value: rand.Intn(2)}
	}
}

func Maj(U []int, n int) (int, int) {
	cnt1 := 0
	for _, v := range U {
		if v == 1 {
			cnt1++
		}
	}
	if cnt1 >= n/2+1 {
		return 1, cnt1
	}
	return 0, len(U) - cnt1
}

func main() {
	var n, s int
	fmt.Scan(&n, &s)

	if n <= 4*s {
		fmt.Println("Trebuie n > 4s")
		return
	}

	vals := make([]int, n)
	for i := range vals {
		fmt.Scan(&vals[i])
	}

	nodes := make([]Node, n)
	for i := 0; i < n; i++ {
		nodes[i] = Node{id: i, x: vals[i], s: s, n: n, inbox: make(chan Message, n)}
	}

	for _, i := range rand.Perm(n)[:s] {
		nodes[i].isByzantine = true
		fmt.Printf("Nodul %d este bizantin\n", i)
	}

	majority := make([]int, n)
	mult := make([]int, n)
	t := 0

	for {
		// Runda 1: fiecare nod isi difuzeaza starea curenta
		for i := 0; i < n; i++ {
			if nodes[i].isByzantine {
				MaliciousBcastState(nodes, i)
			} else {
				Bcast(nodes, i, 1, nodes[i].x)
			}
		}

		// fiecare nod primeste multimea U si calculeaza Majority(t) si mult(t)
		for i := 0; i < n; i++ {
			U := make([]int, 0, n)
			for k := 0; k < n; k++ {
				msg := <-nodes[i].inbox
				U = append(U, msg.Value)
			}
			majority[i], mult[i] = Maj(U, n)
		}

		// Runda 2: king ul fazei t este nodul P_t
		king := t % n
		if nodes[king].isByzantine {
			MaliciousBcastTie(nodes, king)
		} else {
			Bcast(nodes, king, 2, majority[king])
		}

		// actualizarea starii
		for i := 0; i < n; i++ {
			tie := (<-nodes[i].inbox).Value
			if nodes[i].isByzantine {
				continue
			}
			if 2*mult[i] > n+2*s {
				nodes[i].x = majority[i]
			} else {
				nodes[i].x = tie
			}
		}

		if t > s+1 {
			break
		}

		t++
	}

	fmt.Println("Algoritmul s-a oprit la iteratia:", t)

	// toate nodurile corecte trebuie sa aiba aceeasi valoare
	rez := -1
	for i := 0; i < n; i++ {
		if nodes[i].isByzantine {
			continue
		}
		if rez == -1 {
			rez = nodes[i].x
		} else if nodes[i].x != rez {
			fmt.Println("Nu s-a ajuns la consens")
			return
		}
	}
	fmt.Println("Valoarea de consens:", rez)
}
