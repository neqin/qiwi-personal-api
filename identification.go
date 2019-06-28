package qiwi

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

/*
Identification
*/
type Identification struct {
	ID         int                 `json:"id"`
	Type       IdentificationLevel `json:"type"`
	BirthDate  time.Time           `json:"birthDate"`
	FirstName  string              `json:"firstName"`
	MiddleName string              `json:"middleName"`
	LastName   string              `json:"lastName"`
	Passport   string              `json:"passport"`
	Inn        string              `json:"inn"`
	Snils      string              `json:"snils"`
	Oms        string              `json:"oms"`
}

/*
Get identification

https://edge.qiwi.com/identification/v1/persons/wallet/identification
*/
func (qiwi *QiwiPersonalApi) GetIdentification(wallet string) (*Identification, error) {
	resp, err := qiwi.sendRequest(qiwi.apiKey, "GET", fmt.Sprintf("/identification/v1/persons/%s/identification", wallet), nil)
	if err != nil {
		return nil, err
	}

	var res Identification
	err = json.Unmarshal(resp, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (qiwi *QiwiPersonalApi) PostIdentification() error {
	/*not implemented*/
	return errors.New("not implemented")
}
