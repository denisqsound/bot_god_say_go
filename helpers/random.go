package helpers

import (
	"math/rand"
	"time"
)

func Random(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	if min > max {
		return min
	}
	return rand.Intn(max-min) + min
}
