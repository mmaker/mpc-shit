package mpc

import (
	"bytes"
	"math/big"
	"testing"
)

func TestEvaluatePolynomial(t *testing.T) {
	f := []*big.Int{
		big.NewInt(7),
		big.NewInt(4),
		big.NewInt(1),
	}
	mod := big.NewInt(11)
	tests := map[*big.Int]*big.Int{
		big.NewInt(1) : big.NewInt(1),
		big.NewInt(2) : big.NewInt(8),
		big.NewInt(3) : big.NewInt(6),
	}

	for x, expected := range tests {
		got := evaluatePolynomial(f, x, mod)
		if expected.Cmp(got) != 0 {
			t.Errorf("Got %v, expected %v", got, expected)
		}
	}
}

func TestCombine(t *testing.T) {
	// see mpc-book, p. 38
	shares := map[int]*big.Int{
		3 : big.NewInt(6),
		4 : big.NewInt(6),
		5 : big.NewInt(8),
	}
	mod := big.NewInt(11)
	expected := big.NewInt(7)
	got := Combine(shares, mod)
	if expected.Cmp(got) != 0 {
		t.Errorf("Expected %v, got %v", expected, got)
	}

}

func TestSplitAndCombine(t *testing.T) {
	seed := bytes.NewBufferString("forty-two")
	mod := big.NewInt(11)
	secret := big.NewInt(10)
	threshold := 3
	sharesn := 100

	shares := Split(sharesn, threshold, secret, mod, seed)
	if len(shares) != sharesn {
		t.Errorf("Expected %s shares, got %d", sharesn, len(shares))
	}

	var got *big.Int
	var s map[int]*big.Int

	s = map[int]*big.Int{
		1  : shares[0],
		10 : shares[9],
		2  : shares[1],
		5  : shares[4],
		7  : shares[6],
	}
	got = Combine(s, mod)

 	if secret.Cmp(got) != 0 {
 		t.Errorf("Expected secret: %v, got %v", secret, got)
 	}


	s = map[int]*big.Int{
		1  : shares[0],
		2  : shares[1],
		3  : shares[2],
		5  : shares[4],
		6  : shares[5],
	}
	got = Combine(s, mod)
	if secret.Cmp(got) != 0 {
 		t.Errorf("Expected secret: %v, got %v", secret, got)
 	}
}
