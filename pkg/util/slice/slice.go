package slice

import (
	"math"
	"math/rand"
	"strconv"
	"time"
)

// CalcStandardDeviation 配列の標準偏差を計算する
func CalcStandardDeviation(arr []int) float64 {
	sum := Sum(arr)
	len := len(arr)
	ave := float64(sum) / float64(len)

	numerator := 0.0
	for _, v := range arr {
		numerator += math.Pow(float64(v)-ave, 2)
	}
	return math.Sqrt(numerator / float64(len))
}

// CreateTableFilledZero 全ての要素が0のテーブルを生成する
func CreateTableFilledZero(len int) [][]int {
	table := make([][]int, len)
	for i := 0; i < len; i++ {
		table[i] = make([]int, len)
	}
	return table
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

// Shuffle スライスをシャッフルする
func Shuffle(arr []int) {
	rand.Seed(time.Now().UnixNano())
	for i := range arr {
		j := rand.Intn(i + 1)
		arr[i], arr[j] = arr[j], arr[i]
	}
}

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

// Sum int型配列の要素の合計値を算出する
func Sum(arr []int) int {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	return sum
}
