package api

import (
	"math"
	"net/http"

	"local.packages/model"

	"github.com/labstack/echo"
)

// CombinationsInfo 組み分けを生成するために必要な情報
type CombinationsInfo struct {
	// 参加人数
	AllParticipants int
	// 1グループ毎の人数
	ParticipantsInEachGroup int
	// グループ分けの回数
	RepeatCnt int
	// 組み分けの試行回数
	Trials int
}

// SimulationResult 生成結果
type SimulationResult struct {
	ParticipantCombinations [][][]int `json:"participantCombinations"`
	CountTable              [][]int   `json:"countTable"`
	CountTableOfElmNum      []int     `json:"countTableOfElmNum"`
	StandardDeviation       float64   `json:"standardDeviation"`
}

// NewSimulationResult コンストラクタ
func NewSimulationResult(pc *model.ParticipantCombinations, sr *model.ScoreRecord, sd float64, ci *CombinationsInfo) *SimulationResult {

	return &SimulationResult{
		ParticipantCombinations: pc.DevideCombination(ci.ParticipantsInEachGroup),
		CountTable:              sr.CountTable,
		CountTableOfElmNum:      sr.CountNum(ci.RepeatCnt),
		StandardDeviation:       math.Round(sd*100) / 100,
	}
}

// SimulateCombinations 与えられた情報をもとに、最良の組み分けを返却
func SimulateCombinations(c echo.Context) error {
	ci := new(CombinationsInfo)
	if err := c.Bind(ci); err != nil {
		return err
	}

	if ci.AllParticipants%ci.ParticipantsInEachGroup != 0 {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "このプログラムは参加人数が1グループ毎の人数で割り切れない場合に対応していません。割り切れるように数を設定してください。"})
	}

	betterSd := math.MaxFloat64
	var betterPc *model.ParticipantCombinations
	var betterSr *model.ScoreRecord
	for i := 0; i < ci.Trials; i++ {
		pc := model.NewParticipantCombinations(ci.AllParticipants, ci.RepeatCnt)
		sr := model.NewScoreRecord(ci.AllParticipants)

		for _, combination := range pc.Combinations {
			sr.Record(combination, ci.ParticipantsInEachGroup)
		}

		// TODO: 標準偏差が0になる対策、標準偏差が近しい場合は、0が少ない組み分けを採用など...
		tmpSd := sr.CalcStandardDeviation()
		if tmpSd < betterSd {
			betterSd = tmpSd
			betterPc, betterSr = pc, sr
		}
	}

	nsr := NewSimulationResult(betterPc, betterSr, betterSd, ci)
	return c.JSON(http.StatusOK, nsr)
}
