package mpc

import (
	"math/big"
	"os"
	"sync"
)


func CombineParties(parties []*Party) *big.Int {
	mod := parties[0].Mod
	shares := make(map[int]*big.Int, len(parties))
	for _, p := range parties {
		shares[p.Id] = p.Data
	}
	return Combine(shares, mod)
}


func (p *Party) SecretShare(secret *big.Int, t int) {
	urandom, _ := os.Open("/dev/urandom")
	defer urandom.Close()

	shares := Split(len(p.Connections), t, secret, p.Mod, urandom)

	for i, c := range p.Connections {
		go p.SendInt(shares[i], c)
	}
}

func Add(a *big.Int, b *big.Int, parties []*Party, t int) {
	go parties[0].SecretShare(a, t)
	go parties[0].SecretShare(b, t)

	var wg sync.WaitGroup

	wg.Add(len(parties))
	for _, p := range parties {
		go func (p *Party){
			defer wg.Done()

			var z big.Int
			sum := new(big.Int)
			z = <- p.Chan
			sum.Add(sum, &z)
			z = <- p.Chan
			sum.Add(sum, &z)

			p.Data = sum
		}(p)
	}
	wg.Wait()
}

func ConstMul(z *big.Int, parties []*Party, t int) {
	var wg sync.WaitGroup

	for _, p := range parties {
		go func (p *Party) {
			defer wg.Done()
			z.Mul(p.Data, z)
		}(p)
	}
	wg.Wait()
}
