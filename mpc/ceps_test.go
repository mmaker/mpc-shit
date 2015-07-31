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
		t.Errorf("Expected %v, got %v.", sum, tsum)
	}
}


func benchmarkAdd(b *testing.B, actors int) {
	mod := big.NewInt(0x10001)
	parties := SpawnParties(actors, mod)
	first := big.NewInt(1337)
	second := big.NewInt(42)
	threshold := actors / 4
	for i := 0; i != b.N; i++ {
		Add(first, second, parties, threshold)
	}
}


func BenchmarkAdd10(b *testing.B) { benchmarkAdd(b, 10); }
func BenchmarkAdd100(b *testing.B) { benchmarkAdd(b, 100); }
func BenchmarkAdd1000(b *testing.B) { benchmarkAdd(b, 1000); }
