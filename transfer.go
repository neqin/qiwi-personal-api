package qiwi

import (
	"encoding/json"
	"fmt"
	"time"
)

/*
{
  "content": {
      "terms": {
         "commission": {
                "ranges": [{
                     "bound": 0,
                     "fixed": 50.0,
                     "rate": 0.02
                }]
           }
       }
    }
}
*/
type Commissions struct {
	ID      string
	Content struct {
		Terms struct {
			Commission struct {
				Ranges []Ranges
			}
			Limits            []Limit
			Description       string
			Overpayment       bool
			Underpayment      bool
			RepeatablePayment bool
		}
	} `json:"content"`
}

type Ranges struct {
	Bound float64
	Fixed float64
	Rate  float64
}

type Limit struct {
	Currency string
	Min      float64
	Max      float64
}

func (qiwi *QiwiPersonalApi) GetCommissions(provider PaymentProvider) (*Commissions, error) {
	resp, err := qiwi.sendRequest(qiwi.apiKey, "GET", fmt.Sprintf("/sinap/providers/%s/form", provider), nil)
	if err != nil {
		return nil, err
	}

	var res Commissions
	err = json.Unmarshal(resp, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

/*
{
    "id": ":11111111111111",
 	"fields": {
    	    "account": "+79121112233"
  	},
  	"sum": {
    	   "amount": 100,
    	   "currency": "643"
  	},
  	"source": "account_643",
  	"transaction": {
    	   "id": "4969142201",
    	   "state": {
      		"code": "Accepted"
    	    }
  	}
}
*/

type TransferResult struct {
	ID     string
	Fields struct {
		Account string
	}
	Sum struct {
		Amount   float64
		Currency Currency
	}
	Source      string
	Transaction TransferThx
}

type TransferThx struct {
	ID    string
	State struct {
		Code string
	}
}

/*
GetWalletBalance
*/
func (qiwi *QiwiPersonalApi) TransferP2P(to string, currency Currency, amount float64, comment string) (*TransferResult, error) {
	id := time.Now().UnixNano()
	data := map[string]interface{}{
		"id": fmt.Sprint(id),
		"sum": map[string]interface{}{
			"amount":   amount,
			"currency": fmt.Sprint(currency),
		},
		"source": "account_643",
		"paymentMethod": map[string]interface{}{
			"type":      "Account",
			"accountId": "643",
		},
		"comment": comment,
		"fields": map[string]interface{}{
			"account": to,
		},
	}
	resp, err := qiwi.sendRequest(qiwi.apiKey, "POST", "/sinap/api/v2/terms/99/payments", data)
	if err != nil {
		return nil, err
	}

	var res TransferResult
	err = json.Unmarshal(resp, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}
