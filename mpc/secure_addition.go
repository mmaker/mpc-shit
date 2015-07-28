package mpc

import (
	"crypto/rand"
	"io"
	"math/big"
	"sync"
)

func (p *Party) secureAdd(random io.Reader) {
	sum := new(big.Int)

	for _, c := range p.Connections {
		if (c != p.Chan) {
			r, _ := rand.Int(random, p.Mod)
			sum.Add(sum, r)
			sum.Mod(sum, p.Mod)
			go p.SendInt(r, c)
		}
	}
	s := new(big.Int)
	s.Sub(p.Data, sum)
	s.Mod(s, p.Mod)


	for i := 0; i != len(p.Connections)-1; i++ {
		var z big.Int
		z = <- p.Chan
		s.Add(s, &z)
		s.Mod(s, p.Mod)
	}

	go p.SendInt(s, nil)
	sum = new(big.Int)
	for i := 0; i != len(p.Connections); i++ {
		var z big.Int
		z = <- p.Chan
		sum.Add(sum, &z)
		sum.Mod(sum, p.Mod)
	}
	p.Data = sum
}

func SecureAddition(ps []*Party, random io.Reader) {
	var wg sync.WaitGroup
	wg.Add(len(ps))

	for _, p := range ps {
		go func(p *Party) {
			defer wg.Done()
			p.secureAdd(random)
		}(p)
	}
	wg.Wait()
}
