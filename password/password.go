// Package password generates a strong password from a hashed string.
package password

import (
	"hash/fnv"
	"math/rand"
)

// Generator is the interface that creates a fixed length password.
type Generator interface {
	Encrypt() string
}

// New creates a new password.
func New(passwordLength int, passphrase string) Generator {
	g := generator{length: passwordLength, hash: setHash(passphrase)}
	return &g

}

type generator struct {
	hash   int64
	length int
}

func setHash(input string) int64 {
	h := fnv.New64a()
	h.Write([]byte(input))

	value := int64(h.Sum64())
	if value < 0 {
		value = value * -1
	}
	return value
}

func (g *generator) Encrypt() string {
	b := make([]byte, g.length)

	rand.Seed(g.hash)
	for i := 0; i < g.length; i++ {
		b = append(b, byte(33+rand.Intn(93)))
	}
	return string(b)
}
