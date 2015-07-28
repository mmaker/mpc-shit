package mpc

import (
	"crypto/rand"
	"io"
	"math/big"
)

func Split(s int, t int, w *big.Int,  mod *big.Int, random io.Reader) []*big.Int {
	t++;

 	poly := make([]*big.Int, t);
 	poly[0] = w
 	for i := 1; i != t; i++ {
 		poly[i], _ = rand.Int(random, mod)
 	}

 	shares := make([]*big.Int, s)
 	for i := 0; i != s; i++ {
		x := big.NewInt(int64(i+1))
 		shares[i] = evaluatePolynomial(poly, x, mod)
 	}

	return shares
}


func Combine(shares map[int]*big.Int, mod *big.Int) *big.Int {
	secret := new(big.Int)

	for i, share := range shares {
		deltai := big.NewInt(1)
		for j := range shares {
			if (i != j) {
				var z *big.Int
				z = big.NewInt(int64(i-j))
				z.Mod(z, mod)
				z.ModInverse(z, mod)
				deltai.Mul(deltai, z)
				z = big.NewInt(int64(-j))
				deltai.Mul(deltai, z)
				deltai.Mod(deltai, mod)
			}
		}
		deltai.Mul(deltai, share)
		secret.Add(secret, deltai)
		secret.Mod(secret, mod)
	}
	return secret
}
