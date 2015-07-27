package mpc

import (
	"fmt"
	"math/rand"
	"sync"
)
const MOD = 0x1001

type Player struct {
	Name string
	Chan chan int
	Connections []chan int
	Secret int
	Result int
}

func (p *Player) Send(what int, whom chan int) {
	if whom != nil {
		whom <- what
	} else {
		for _, c := range p.Connections {
			p.Send(what, c)
		}
	}
}


func (p *Player) shareSecret() {
	sum := 0
	for _, c := range p.Connections {
		r := rand.Intn(MOD)
		sum = (sum + r) % MOD
		go p.Send(r, c)
	}

	s := (p.Secret - sum) % MOD
	for i := 0; i != len(p.Connections); i++ {
		s = (s + <- p.Chan) % MOD
	}

	go p.Send(s, nil)
	p.Result = s
	for i := 0; i != len(p.Connections); i++ {
		p.Result = (p.Result + <- p.Chan) % MOD
	}
}

func SecureAddition(ps []*Player) {
	var wg sync.WaitGroup
	wg.Add(len(ps))

	for _, p := range ps {
		go func(p *Player) {
			defer wg.Done()
			p.shareSecret()
		}(p)
	}
	wg.Wait()
}


func main() {
	p1 := Player {
		Name : "p1",
		Chan : make(chan int),
		Secret : 1,
	}
	p2 := Player {
		Name : "p2",
		Chan : make(chan int),
		Secret : 1,
	}

	p3 := Player {
		Name : "p3",
		Chan : make(chan int),
		Secret : 0,
	}


	p1.Connections = []chan int{p2.Chan, p3.Chan}
	p2.Connections = []chan int{p1.Chan, p3.Chan}
	p3.Connections = []chan int{p1.Chan, p2.Chan}

	SecureAddition([]*Player{&p1, &p2, &p3})
	fmt.Println(p1.Result)
}
