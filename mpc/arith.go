package mpc

import "math/big"

func evaluatePolynomial(coefficients []*big.Int, x *big.Int, mod *big.Int) *big.Int {
	term := new(big.Int)
	z := new(big.Int)

 	for i, c := range coefficients {
		term = big.NewInt(int64(i))
		term.Exp(x, term, mod)
		term.Mul(term, c)
		z.Add(z, term)
		z.Mod(z, mod)
 	}
	return z
}
