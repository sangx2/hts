package hts

import (
	"encoding/json"
	"github.com/sangx2/ebest/res"
	"io"
	"strings"
)

type KOSPI struct {
	H1 res.H1OutBlock `json:"호가잔량,omitempty"`
	S3 res.S3OutBlock `json:"체결,omitempty"`
	K1 res.K1OutBlock `json:"거래원,omitempty"`
}

type KOSDAQ struct {
	HA res.HAOutBlock `json:"호가잔량,omitempty"`
	K3 res.K3OutBlock `json:"체결,omitempty"`
	OK res.OKOutBlock `json:"거래원,omitempty"`
}

type Stock struct {
	res.T8436OutBlock `json:"기본정보"`

	*KOSPI  `json:"KOSPI,omitempty"`
	*KOSDAQ `json:"KOSDAQ,omitempty"`

	CSPAT00600OutBlock1 map[string]res.CSPAT00600OutBlock1 `json:"현물정상주문1"`
	CSPAT00600OutBlock2 map[string]res.CSPAT00600OutBlock2 `json:"현물정상주문2"`

	CSPAT00700OutBlock1 map[string]res.CSPAT00700OutBlock1 `json:"현물정정주문1"`
	CSPAT00700OutBlock2 map[string]res.CSPAT00700OutBlock2 `json:"현물정정주문2"`

	CSPAT00800OutBlock1 map[string]res.CSPAT00800OutBlock1 `json:"현물취소주문1"`
	CSPAT00800OutBlock2 map[string]res.CSPAT00800OutBlock2 `json:"현물취소주문2"`

	SC0OutBlock map[string]res.SC0OutBlock `json:"접수"`
	SC1OutBlock map[string]res.SC1OutBlock `json:"체결"`
	SC2OutBlock map[string]res.SC2OutBlock `json:"정정"`
	SC3OutBlock map[string]res.SC3OutBlock `json:"취소"`
	SC4OutBlock map[string]res.SC4OutBlock `json:"거부"`
}

func NewStock(t8436OutBlock res.T8436OutBlock) *Stock {
	s := &Stock{
		T8436OutBlock:       t8436OutBlock,
		CSPAT00600OutBlock1: make(map[string]res.CSPAT00600OutBlock1),
		CSPAT00600OutBlock2: make(map[string]res.CSPAT00600OutBlock2),
		CSPAT00700OutBlock1: make(map[string]res.CSPAT00700OutBlock1),
		CSPAT00700OutBlock2: make(map[string]res.CSPAT00700OutBlock2),
		CSPAT00800OutBlock1: make(map[string]res.CSPAT00800OutBlock1),
		CSPAT00800OutBlock2: make(map[string]res.CSPAT00800OutBlock2),

		SC0OutBlock: make(map[string]res.SC0OutBlock),
		SC1OutBlock: make(map[string]res.SC1OutBlock),
		SC2OutBlock: make(map[string]res.SC2OutBlock),
		SC3OutBlock: make(map[string]res.SC3OutBlock),
		SC4OutBlock: make(map[string]res.SC4OutBlock),
	}
	if strings.Compare(s.T8436OutBlock.Gubun, "1") == 0 {
		s.KOSPI = new(KOSPI)
	} else {
		s.KOSDAQ = new(KOSDAQ)
	}
	return s
}

func StockFromJson(data io.Reader) *Stock {
	var stock *Stock

	json.NewDecoder(data).Decode(&stock)

	return stock
}

func (s *Stock) ToJson() []byte {
	b, e := json.Marshal(s)
	if e != nil {
		return nil
	}

	return b
}
