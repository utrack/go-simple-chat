package clientBot

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func random(min, max int) int {
	return rand.Intn(max-min) + min
}
