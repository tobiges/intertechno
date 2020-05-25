package intertechno

import (
	"math/rand"
	"time"
)

var r = rand.New(rand.NewSource(time.Now().UnixNano()))

// GenerateAddress returns pseudo random address
func GenerateAddress() int {
	return r.Int() & (1<<addressBits - 1)
}

// GenerateUnit returns pseudo random unit
func GenerateUnit() int {
	return r.Int() & (1<<unitBits - 1)
}
