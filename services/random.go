package services

import (
	"math/rand"
	"time"
)

func Randomize() int {
	rand.Seed(time.Now().UnixNano())
	num := rand.Intn(100-1) + 1

	return num
}
