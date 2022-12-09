package model

import (
	"strings"

	"github.com/yagi_eng/group_shuffle_gui/util/slice"
)

// ParticipantCombinations グループ分けを管理する型
type ParticipantCombinations struct {
	Combinations [][]int
}

// NewParticipantCombinations コンストラクタ
func NewParticipantCombinations(allParticipants int, repeatCnt int) *ParticipantCombinations {
	combination := []int{}
	for i := 0; i < allParticipants; i++ {
		combination = append(combination, i)
	}

	combinations := make([][]int, repeatCnt)
	for i := 0; i < repeatCnt; i++ {
		slice.Shuffle(combination)
		combinations[i] = make([]int, len(combination))
		copy(combinations[i], combination)
	}

	return &ParticipantCombinations{Combinations: combinations}
}

// CreateCombinationsForFront フロント表示用に文字列（数字or名前）に変換する
func (pc *ParticipantCombinations) CreateCombinationsForFront(names string, participantsInEachGroup int) [][][]string {
	combinationsStr := [][]string{}
	if names == "" {
		combinationsStr = slice.Itoa(pc.Combinations)
	} else {
		combinationsStr = pc.replaceNumWithName(names)
	}
	return pc.devideCombination(combinationsStr, participantsInEachGroup)
}

// replaceNumWithName 数字を人の名前に置換する
func (pc *ParticipantCombinations) replaceNumWithName(names string) [][]string {
	namesArr := strings.Split(names, ",")
	combinationsStr := [][]string{}

	for _, combination := range pc.Combinations {
		combinationStr := []string{}
		for _, val := range combination {
			combinationStr = append(combinationStr, namesArr[val])
		}
		combinationsStr = append(combinationsStr, combinationStr)
	}

	return combinationsStr
}

// devideCombination シミュレーション結果表示用に各組み合わせをグループ毎に分割する
func (pc *ParticipantCombinations) devideCombination(combinationsStr [][]string, participantsInEachGroup int) [][][]string {
	devideCombinations := [][][]string{}
	for _, combination := range combinationsStr {
		devideCombinations = append(devideCombinations, slice.DevideArrStr(combination, participantsInEachGroup))
	}
	return devideCombinations
}
