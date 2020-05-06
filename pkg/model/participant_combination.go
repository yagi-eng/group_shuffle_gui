package model

import (
	"strings"

	array "local.packages/pkg/util/array"
)

// ParticipantCombinations 参加者の組み合わせを管理する型
type ParticipantCombinations struct {
	Combinations    [][]int
	CombinationsStr [][]string
}

// NewParticipantCombinations コンストラクタ
func NewParticipantCombinations(allParticipants int, repeatCnt int) *ParticipantCombinations {
	combinations := make([][]int, repeatCnt)

	combination := make([]int, allParticipants)
	for i := range combination {
		combination[i] = i
	}

	for i := 0; i < repeatCnt; i++ {
		array.Shuffle(combination)
		combinations[i] = make([]int, len(combination))
		copy(combinations[i], combination)
	}

	return &ParticipantCombinations{Combinations: combinations}
}

// CreateCombinationsStr string型のCombinationsを作成する
func (pc *ParticipantCombinations) CreateCombinationsStr() {
	pc.CombinationsStr = array.Itoa(pc.Combinations)
}

// ReplaceNumWithName 数字を人の名前に置換する
func (pc *ParticipantCombinations) ReplaceNumWithName(names string) {
	namesArr := strings.Split(names, ",")
	for i, combination := range pc.Combinations {
		for j, elm := range combination {
			pc.CombinationsStr[i][j] = namesArr[elm]
		}
	}
}

// DevideCombination シミュレーション結果表示用に各組み合わせをグループ毎に分割する
func (pc *ParticipantCombinations) DevideCombination(participantsInEachGroup int) [][][]string {
	devideCombinations := [][][]string{}
	for _, combination := range pc.CombinationsStr {
		devideCombinations = append(devideCombinations, array.SliceArrStr(combination, participantsInEachGroup))
	}
	return devideCombinations
}
