package handlers

import (
	"math"
	"net/http"

	"local.packages/models"

	"github.com/labstack/echo"
)

// CombinationsInfo 組み分けを生成するために必要な情報
type CombinationsInfo struct {
	// 参加者数
	AllParticipants int
	// 1グループ毎の人数
	ParticipantsInEachGroup int
	// 組み分けの繰り返し回数
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
func NewSimulationResult(pc *models.ParticipantCombinations, sr *models.ScoreRecord, sd float64, ci *CombinationsInfo) *SimulationResult {

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
		return c.JSON(http.StatusOK, map[string]string{"message": "このプログラムは参加人数がグループ数で割り切れない場合に対応していません。割り切れるように数を設定してください。"})
	}

	betterSd := math.MaxFloat64
	var betterPc *models.ParticipantCombinations
	var betterSr *models.ScoreRecord
	for i := 0; i < ci.Trials; i++ {
		pc := models.NewParticipantCombinations(ci.AllParticipants, ci.RepeatCnt)
		sr := models.NewScoreRecord(ci.AllParticipants)

		for _, combination := range pc.Combinations {
			sr.Record(combination, ci.ParticipantsInEachGroup)
		}

		// TODO: 標準偏差が0になる対策、標準偏差が近しい場合は、0が多い組み分けを採用など...
		tmpSd := sr.CalcStandardDeviation()
		if tmpSd < betterSd {
			betterSd = tmpSd
			betterPc, betterSr = pc, sr
		}
	}

	nsr := NewSimulationResult(betterPc, betterSr, betterSd, ci)
	return c.JSON(http.StatusOK, nsr)
}
