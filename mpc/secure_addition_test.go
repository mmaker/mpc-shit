package mpc

import (
	"math/big"
	"os"
	"testing"
)

func TestSecureAddition(t *testing.T) {
	urandom, _ := os.Open("/dev/urandom")
	parties := SpawnParties(3, big.NewInt(11))
	parties[0].Data = big.NewInt(10)
	parties[1].Data = big.NewInt(0)
	parties[2].Data = big.NewInt(1)
	expected := big.NewInt(0)

	SecureAddition(parties, urandom)
	if parties[0].Data.Cmp(parties[1].Data) != 0 ||
		parties[1].Data.Cmp(parties[2].Data) != 0 {
		t.Errorf("Expecting all parties to share the same sum.")
	}

	got := parties[0].Data
	if got.Cmp(expected) != 0 {
		t.Errorf("Expected %v, got %v", expected, got)
	}

}
