package qiwi

import (
	"encoding/json"
	"fmt"
	"time"
)

/*
PaymentsHistory
{"data":[{}],
  "nextTxnId":9001,
  "nextTxnDate":"2017-01-31T15:24:10+03:00"
}
*/
type PaymentsHistory struct {
	Data        []PaymentsHistoryData `json:"data"`
	NextTxnID   int                   `json:"nextTxnId"`
	NextTxnDate time.Time             `json:"nextTxnDate"`
}

/*
PaymentsHistoryData
{
    "txnId":9309,
    "personId":79112223344,
    "date":"2017-01-21T11:41:07+03:00",
    "errorCode":0,
    "error":null,
    "status":"SUCCESS",
    "type":"OUT",
    "statusText":"Успешно",
    "trmTxnId":"1489826461807",
    "account":"0003***",
    "sum":{
        "amount":70,
        "currency":"RUB"
        },
    "commission":{
        "amount":0,
        "currency":"RUB"
        },
    "total":{
        "amount":70,
        "currency":"RUB"
        },
    "provider":{
                       "id":26476,
                       "shortName":"Yandex.Money",
                       "longName":"ООО НКО \"Яндекс.Деньги\"",
                       "logoUrl":"https://static.qiwi.com/img/providers/logoBig/26476_l.png",
                       "description":"Яндекс.Деньги",
                       "keys":"***",
                       "siteUrl":null
                      },
    "comment":null,
    "currencyRate":1,
    "extras":null,
    "chequeReady":true,
    "bankDocumentAvailable":false,
    "bankDocumentReady":false,
                "repeatPaymentEnabled":false
	}
*/
type PaymentsHistoryData struct {
	TxnID                 int `json:"txnId"`
	PersonID              int `json:"personID"`
	Date                  time.Time
	ErrorCode             int
	Error                 string
	Status                PaymentStatus
	Type                  PaymentHistoryOperation `json:"type"`
	StatusText            string
	TrmTxnID              string `json:"trmTxnId"`
	Account               string
	Sum                   PaymentAmount
	Commission            PaymentAmount
	Total                 PaymentAmount
	Provider              PaymentsHistoryProvider
	Comment               string
	CurrencyRate          float64
	Extras                interface{}
	ChequeReady           bool
	BankDocumentAvailable bool
	BankDocumentReady     bool
	RepeatPaymentEnabled  bool
}

type PaymentAmount struct {
	Amount   float32
	Currency Currency
}

type PaymentsHistoryProvider struct {
	ID          int `json:"id"`
	ShortName   string
	LongName    string
	LogoURL     string `json:"LogoUrl"`
	Description string
	Keys        string
	SiteURL     string `json:"siteUrl"`
}

/*
GetPaymentsHistory

История платежей и пополнений вашего кошелька.
*/
func (qiwi *QiwiPersonalApi) GetPaymentsHistory(wallet string, rows int) (*PaymentsHistory, error) {
	resp, err := qiwi.sendRequest(qiwi.apiKey, "GET", fmt.Sprintf("/payment-history/v1/persons/%s/payments?rows=%d", wallet, rows), nil)
	if err != nil {
		return nil, err
	}

	var res PaymentsHistory
	err = json.Unmarshal(resp, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

/*
{
 "incomingTotal":[
  {
  "amount":3500,
  "currency":"RUB"
  }],
 "outgoingTotal":[
  {
  "amount":3497.5,
  "currency":"RUB"
  }]
}
*/
type PaymentsStats struct {
	IncomingTotal []PaymentAmount
	OutgoingTotal []PaymentAmount
}

/*
GetPaymentsStats

PH_OPERATION_ALL 		= "ALL"
PH_OPERATION_IN 		= "IN"
PH_OPERATION_OUT 		= "OUT"
PH_OPERATION_QIWI_CARD 	= "QIWI_CARD"
*/
func (qiwi *QiwiPersonalApi) GetPaymentsStats(wallet string, op PaymentHistoryOperation) (*PaymentsStats, error) {
	resp, err := qiwi.sendRequest(qiwi.apiKey, "GET", fmt.Sprintf("/payment-history/v1/persons/%s/payments/total?operation=%s", wallet, op), nil)
	if err != nil {
		return nil, err
	}

	var res PaymentsStats
	err = json.Unmarshal(resp, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}
