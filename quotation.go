package hts

import (
	"encoding/json"
	"github.com/sangx2/ebest/res"
	"io"
)

type Quotation struct {
	StockName string

	T1305OutBlock1s map[string]res.T1305OutBlock1 `json:"기간별 주가"` // map[date]
}

func NewQuotation(stockName string) *Quotation {
	return &Quotation{
		StockName:       stockName,
		T1305OutBlock1s: make(map[string]res.T1305OutBlock1),
	}
}

func QuotationFromJson(data io.Reader) *Quotation {
	var quotation *Quotation

	json.NewDecoder(data).Decode(&quotation)

	return quotation
}

func (q *Quotation) ToJson() []byte {
	b, e := json.Marshal(q)
	if e != nil {
		return nil
	}

	return b
}
