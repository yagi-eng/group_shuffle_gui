package pkg

import (
	"math/rand"
	"strconv"
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

// SliceArrStr スライスを分割する
func SliceArrStr(arr []string, lenOfEachSlice int) [][]string {
	arrs := [][]string{}
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

// Itoa int型の2次元配列をstring型の2次元配列に置き換える
func Itoa(tableOfInt [][]int) [][]string {
	tableOfStr := [][]string{}
	for _, arrOfInt := range tableOfInt {
		arrOfStr := []string{}
		for _, val := range arrOfInt {
			arrOfStr = append(arrOfStr, strconv.Itoa(val))
		}
		tableOfStr = append(tableOfStr, arrOfStr)
	}
	return tableOfStr
}
