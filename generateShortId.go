package main

import (
	"crypto/rand"
	"log"
	"math/big"
)

var (
	lettersRune = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
)

func generateShortId() string {
	b := make([]rune, 6)
	for i := range b {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(lettersRune))))
		if err != nil {
			log.Fatal(err)
		}
		b[i] = lettersRune[num.Int64()]
	}

	return string(b)
}
