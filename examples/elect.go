package main


import (
	"fmt"
	"math/big"
	"os"

	"github.com/mmaker/mpc-shit/mpc"
)

func main() {
	urandom, _ := os.Open("/dev/urandom")

	n := 3
	parties := mpc.SpawnParties(n, big.NewInt(1337))
	parties[0].Data = big.NewInt(1)
	parties[1].Data = big.NewInt(0)
	parties[2].Data = big.NewInt(0)

	mpc.SecureAddition(parties, urandom)
	sum := parties[0].Data
	majority := big.NewInt(int64(n))
	majority.Rsh(majority, 1)
	if (sum.Cmp(majority) > 0) {
		fmt.Println("Won")
	} else {
		fmt.Println("Lost")
	}
}
