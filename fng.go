package hts

import (
	"encoding/json"
	"github.com/sangx2/ebest/res"
	"io"
)

type FNG struct {
	Name string `json:"종목명"`
	// Stock 요약
	res.T3320OutBlock  `json:"기업기본정보"`
	res.T3320OutBlock1 `json:"기업재무정보"`
}

func FNGFromJson(data io.Reader) *FNG {
	var fng *FNG

	json.NewDecoder(data).Decode(&fng)

	return fng
}

func (p *FNG) ToJson() []byte {
	b, e := json.Marshal(p)
	if e != nil {
		return nil
	}

	return b
}
