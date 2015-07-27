package mpc

import (
	"math/big"
)

type Party struct {
	Id int
	Chan chan big.Int
	Connections []chan big.Int
	Mod *big.Int
	Data *big.Int
}

func (p *Party) SendInt(what *big.Int, whom chan big.Int) {
	if whom != nil {
		whom <- *what
	} else {
		for _, c := range p.Connections {
			p.SendInt(what, c)
		}
	}
}


func SpawnParties(n int, mod *big.Int) []*Party {
	parties := make([]*Party, n)
	connections := make([]chan big.Int, n)

	for i := 0; i != n; i++ {
		p := new(Party)
		p.Id = i+1
		p.Chan = make(chan big.Int)
		p.Connections = connections
		p.Mod = mod

		parties[i] = p
		connections[i] = p.Chan
	}
	return parties
}
