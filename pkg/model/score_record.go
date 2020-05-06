package model

import (
	"math"

	array "local.packages/pkg/util/array"
)

// ScoreRecord スコアと同席回数を管理する型
type ScoreRecord struct {
	// 同席回数をカウントするテーブル(参加者数*参加者数)
	// 例) countTable[3][6] = 2　⇒　参加者3と参加者6の同席回数は2
	CountTable [][]int
	// スコア
	// 参加者毎に各回で同席した人との同席回数を加算していった合計
	Scores []int
}

// NewScoreRecord コンストラクタ
func NewScoreRecord(len int) *ScoreRecord {
	scores := make([]int, len)
	countTable := createTableFilledZero(len)
	return &ScoreRecord{Scores: scores, CountTable: countTable}
}

// 全ての要素が0のテーブルを生成する
func createTableFilledZero(len int) [][]int {
	table := make([][]int, len)
	for i := 0; i < len; i++ {
		table[i] = make([]int, len)
	}
	return table
}

// Record 同席回数とスコアを記録する
func (sr *ScoreRecord) Record(participants []int, participantsInEachGroup int) {
	groups := array.SliceArr(participants, participantsInEachGroup)
	for _, group := range groups {
		sr.recordEachGroup(group)
	}
}

func (sr *ScoreRecord) recordEachGroup(group []int) {
	len := len(group)
	for i := 0; i < len; i++ {
		for j := 0; j < len; j++ {
			if i == j {
				continue
			}

			x, y := group[i], group[j]
			// 同席回数を記録
			sr.CountTable[x][y]++
			// スコアを記録
			sr.Scores[x] += sr.CountTable[x][y]
		}
	}
}

// CalcStandardDeviation スコアの標準偏差を計算する
func (sr *ScoreRecord) CalcStandardDeviation() float64 {
	sum := 0
	for _, score := range sr.Scores {
		sum += score
	}
	len := len(sr.Scores)
	ave := float64(sum) / float64(len)

	numerator := 0.0
	for _, v := range sr.Scores {
		numerator += math.Pow(float64(v)-ave, 2)
	}
	return math.Sqrt(numerator / float64(len))
}

// CountNum 同席回数を集計する
func (sr *ScoreRecord) CountNum(repeatCnt int) []int {
	cnt := make([]int, repeatCnt+1)
	for _, row := range sr.CountTable {
		for _, v := range row {
			cnt[v]++
		}
	}
	return cnt
}
