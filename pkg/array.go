package pkg

import (
	"math/rand"
	"time"
)

// SliceArr スライスを分割する
func SliceArr(arr []int, lenOfEachSlice int) [][]int {
	arrs := [][]int{}
	len := len(arr)

	for i := 0; i < len; i += lenOfEachSlice {
		end := i + lenOfEachSlice
		if len < end {
			end = len
		}
		arrs = append(arrs, arr[i:end])
	}

	return arrs
}

// Shuffle スライスをシャッフルする
func Shuffle(arr []int) {
	rand.Seed(time.Now().UnixNano())
	for i := range arr {
		j := rand.Intn(i + 1)
		arr[i], arr[j] = arr[j], arr[i]
	}
}
