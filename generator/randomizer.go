package generator

import (
	"math/rand"
	"time"
)

type Randomizer struct {
}

var seededRand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func (r *Randomizer) GetRandomInt(n int) int {
	return seededRand.Intn(n)
}

