package hts

import (
	"encoding/json"
	"github.com/sangx2/ebest/res"
	"io"
)

// Account : 계좌 정보
type Account struct {
	Number     string `json:"계좌번호"`
	Name       string `json:"이름"`
	DetailName string `json:"상세"`
	NickName   string `json:"별명"`

	// [ebest]
	//
	// 현물계좌 예수금/주문가능금액/총평가 조회(API)
	Asset res.CSPAQ12200OutBlock2 `json:"자산"`
	//
	// 주식잔고2
	BalanceTotal res.T0424OutBlock    `json:"잔고종합"`
	Balances     []res.T0424OutBlock1 `json:"잔고"`
}

func NewAccount(number, name, detailName, nickName string) *Account {
	return &Account{
		Number:     number,
		Name:       name,
		DetailName: detailName,
		NickName:   nickName,
	}
}

func AccountFromJson(data io.Reader) *Account {
	var account *Account

	json.NewDecoder(data).Decode(&account)

	return account
}
