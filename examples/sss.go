package main

import (
	"fmt"
	"flag"
	"math/big"
	"os"

	"github.com/mmaker/mpc-shit/mpc"

)
const modHex = "FB82E4EFA06A7F5D1CD10395F2EEE60CC9DE2F932C295DFAAC173C8F686B655B"

func Split() {
	fmt.Print("Shares: ")
	var s int
	fmt.Scanf("%d", &s)

	fmt.Print("Threshold: ")
	var t int
	fmt.Scanf("%d", &t)

	fmt.Print("Secret: ")
	n := new(big.Int)
	fmt.Scanf("%X", n)

	urandom, _ := os.Open("/dev/urandom")
	mod := new(big.Int)
	fmt.Sscanf(modHex, "%X", mod)
	shares := mpc.Split(s, t, n, mod, urandom)

	fmt.Printf("Shares for %X:\n", n)
	for i, share := range(shares) {
		fmt.Printf("%d:%X\n", i+1, share)
	}
}

func Combine() {
	var i int
	share := new(big.Int)
	mod := new(big.Int)
	fmt.Sscanf(modHex, "%X", mod)

	nextShare := func () bool {
		read, _ := fmt.Scanf("%d:%X\n", &i, share)
		return read == 2
	}

	shares := make(map[int]*big.Int)
	for nextShare() {
		shares[i] = new(big.Int).Set(share)
	}
	fmt.Printf("%X", mpc.Combine(shares, mod))
}


func main() {
	var combine  = flag.Bool("combine", false, "Recombines the given secret")
	var split = flag.Bool("split", false, "Shares a secret")

	flag.Parse()
	if *combine {
		Combine()
	} else if *split {
		Split()
	} else {
		flag.Usage()
	}
}
