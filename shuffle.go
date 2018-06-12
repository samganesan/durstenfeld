// implements the Durstenfeld or the Fisher-Yates Shuffle
package durstenfeld

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func ShuffleTheSlice(slice []interface{}) {
	// the shuffle interface... I will write int stings and floats
	// for implementations

	l := len(slice)

	/*
		iterate from the last index to the first index and do the swap

		-- To shuffle an array a of n elements (indices 0..n-1):
		for i from n−1 downto 1 do
		j ← random integer such that 0 ≤ j ≤ i
		exchange a[j] and a[i]

	*/
	for i := l - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		slice[i], slice[j] = slice[j], slice[i]
	}
}

func ShuffleTheSliceInt(slice []int) {
	l := len(slice)

	for i := l - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		slice[i], slice[j] = slice[j], slice[i]
	}
}

func ShuffleTheSliceString(slice []string) {
	l := len(slice)

	for i := l - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		slice[i], slice[j] = slice[j], slice[i]
	}
}
