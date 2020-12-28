package hts

import (
	"encoding/json"
	"io"
)

type Portfolio struct {
	StartDate string `json:"시작일"`
	EndDate   string `json:"종료일"`

	StartAsset float64 `json:"초기 자산"`
	EndAsset   float64 `json:"최종 자산"`

	CAGR              float64 `json:"연복리 수익률"`
	MDD               float64 `json:"최대 낙폭"`
	SharpeRatio       float64 `json:"샤프지수"`
	StandardDeviation float64 `json:"표준편차"`

	ProfitAndLoss struct {
		Best  float64 `json:"최고의 해"`
		Worst float64 `json:"최악의 해"`
	} `json:"손익정보"`
}

func NewPortfolio() *Portfolio {
	return &Portfolio{}
}

func PortfolioFromJson(data io.Reader) *Portfolio {
	var portfolio *Portfolio

	json.NewDecoder(data).Decode(&portfolio)

	return portfolio
}

func (p *Portfolio) ToJson() []byte {
	b, e := json.Marshal(p)
	if e != nil {
		return nil
	}

	return b
}
