package qiwi

import (
	"encoding/json"
)

type WalletBalance struct {
	Accounts []WalletBalanceAccount
}

/*
{
            "alias": "mc_beeline_rub",
            "fsAlias": "qb_mc_beeline",
            "title": "MC",
            "type": {
                "id": "MC",
                "title": "Счет мобильного кошелька"
            },
            "hasBalance": false,
            "balance": null,
            "currency": 643
		},
*/
type WalletBalanceAccount struct {
	Alias   string
	FsAlias string
	Title   string
	Type    struct {
		ID    string
		Title string
	}
	HasBalance bool
	Balance    PaymentAmount
	currency   Currency
}

/*
GetWalletBalance
*/
func (qiwi *QiwiPersonalApi) GetWalletBalance() (*WalletBalance, error) {
	resp, err := qiwi.sendRequest(qiwi.apiKey, "GET", "/funding-sources/v1/accounts/current", nil)
	if err != nil {
		return nil, err
	}

	var res WalletBalance
	err = json.Unmarshal(resp, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}
