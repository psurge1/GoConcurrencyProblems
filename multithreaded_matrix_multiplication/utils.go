package multithreadedmatrixmultiplication

import (
	"math/rand/v2"
)

func RandInRange(min int, max int) int {
	if max <= min {
		return 0
	}
	return rand.IntN(max-min) + min
}
