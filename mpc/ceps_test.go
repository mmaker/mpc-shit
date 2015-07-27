package mpc

import (
	"math/big"
	"testing"
)


func TestAdd(t *testing.T) {
	var mod, a, b, sum, tsum *big.Int
	var parties []*Party

	mod = big.NewInt(97)
	parties = SpawnParties(100, mod)
	a = big.NewInt(2)
	b = big.NewInt(1)
	sum = big.NewInt(3)
	Add(a, b, parties, 5)
	tsum = CombineParties(parties[1:8])
	if tsum.Cmp(sum) != 0 {
		t.Errorf("Expected %v, got %v", sum, tsum)
	}
}
