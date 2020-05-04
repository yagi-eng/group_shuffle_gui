package model

import "local.packages/pkg"

// ParticipantCombinations 参加者の組み合わせを管理する型
type ParticipantCombinations struct {
	Combinations [][]int
}

// NewParticipantCombinations コンストラクタ
func NewParticipantCombinations(allParticipants int, repeatCnt int) *ParticipantCombinations {
	combinations := make([][]int, repeatCnt)

	combination := make([]int, allParticipants)
	for i := range combination {
		combination[i] = i
	}

	for i := 0; i < repeatCnt; i++ {
		pkg.Shuffle(combination)
		combinations[i] = make([]int, len(combination))
		copy(combinations[i], combination)
	}

	return &ParticipantCombinations{Combinations: combinations}
}

// DevideCombination シミュレーション結果表示用に各組み合わせをグループ毎に分割する
func (pc *ParticipantCombinations) DevideCombination(participantsInEachGroup int) [][][]int {
	devideCombinations := [][][]int{}
	for _, combination := range pc.Combinations {
		devideCombinations = append(devideCombinations, pkg.SliceArr(combination, participantsInEachGroup))
	}
	return devideCombinations
}
