package dice

import (
	"crypto/rand"
	"log"
	"math/big"
)

func Die(sides int64) int {

	x, err := rand.Int(rand.Reader, big.NewInt(sides))

	if err != nil {
		log.Fatal(err)
	}

	return int(x.Int64()) + 1
}
