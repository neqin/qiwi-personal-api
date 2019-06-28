package qiwi

import (
	"encoding/json"
	"time"
)

/*
PersonProfile
{
  "authInfo": {},
  "contractInfo": {},
  "userInfo": {}
}
*/
type PersonProfile struct {
	AuthInfo     PersonProfileAuthInfo     `json:"authInfo"`
	ContractInfo PersonProfileContractInfo `json:"contractInfo"`
	UserInfo     PersonProfileUserInfo     `json:"userInfo"`
}

/*
PersonProfileAuthInfo
{
    "boundEmail": "m@ya.ru",
    "ip": "81.210.201.22",
    "lastLoginDate": "2017-07-27T06:51:06.099Z",
    "mobilePinInfo": {},
    "passInfo": {},
    "personId": 79683851815,
    "pinInfo": {
      "pinUsed": true
    },
    "registrationDate": "2017-01-07T16:51:06.100Z"
  }
*/
type PersonProfileAuthInfo struct {
	BoundEmail    string                     `json:"boundEmail"`
	IP            string                     `json:"ip"`
	LastLoginDate time.Time                  `json:"lastLoginDate"`
	MobilePinInfo PersonProfileMobilePinInfo `json:"mobilePinInfo"`
	PassInfo      PersonProfilePassInfo      `json:"passInfo"`
	PersonID      int                        `json:"personId"`
	PinInfo       struct {
		PinUsed bool `json:"pinUsed"`
	} `json:"pinInfo"`
	RegistrationDate time.Time `json:"registrationDate"`
}

/*
PersonProfileMobilePinInfo
{
  "lastMobilePinChange": "2017-07-13T11:22:06.099Z",
  "mobilePinUsed": true,
  "nextMobilePinChange": "2017-11-27T06:51:06.099Z"
}
*/
type PersonProfileMobilePinInfo struct {
	LastMobilePinChange time.Time `json:"lastMobilePinChange"`
	MobilePinUsed       bool      `json:"mobilePinUsed"`
	NextMobilePinChange time.Time `json:"nextMobilePinChange"`
}

/*
PersonProfilePassInfo
{
  "lastPassChange": "2017-07-21T09:25:06.099Z",
  "nextPassChange": "2017-08-21T09:25:06.099Z",
  "passwordUsed": true
}
*/
type PersonProfilePassInfo struct {
	LastPassChange time.Time `json:"lastPassChange"`
	NextPassChange time.Time `json:"nextPassChange"`
	PasswordUsed   bool      `json:"passwordUsed"`
}

/*
PersonProfileContractInfo
{
    "blocked": false,
    "contractId": 79683851815,
    "creationDate": "2017-01-07T16:51:06.100Z",
    "features": [
      ...
    ],
    "identificationInfo": [
      {
        "bankAlias": "QIWI",
        "identificationLevel": "SIMPLE"
      }
    ]

*/
type PersonProfileContractInfo struct {
	Blocked            bool
	ContractID         int           `json:"contractId"`
	CreationDate       time.Time     `json:"creationDate"`
	Features           []interface{} `json:"features"`
	IdentificationInfo []struct {
		BankAlias           string              `json:"bankAlias"`
		IdentificationLevel IdentificationLevel `json:"identificationLevel"`
	} `json:"identificationInfo"`
}

/*
PersonProfileUserInfo
{
  "defaultPayCurrency": 643,
  "defaultPaySource": 7,
  "email": null,
  "firstTxnId": 10807097143,
  "language": "string",
  "operator": "Beeline",
  "phoneHash": "lgsco87234f0287",
  "promoEnabled": null
}
*/
type PersonProfileUserInfo struct {
	DefaultPayCurrency Currency `json:"defaultPayCurrency"`
	DefaultPaySource   int      `json:"defaultPaySource"`
	Email              string   `json:"email"`
	FirstTxnID         int64    `json:"firstTxnId"`
	Language           string   `json:"language"`
	Operator           string   `json:"operator"`
	PhoneHash          string   `json:"phoneHash"`
	PromoEnabled       string   `json:"promoEnabled"`
}

/*
Person’s Profile

https://edge.qiwi.com/person-profile/v1/profile/current?<parameters>
*/
func (qiwi *QiwiPersonalApi) GetPersonProfile(param string) (*PersonProfile, error) {
	resp, err := qiwi.sendRequest(qiwi.apiKey, "GET", "/person-profile/v1/profile/current?"+param, nil)
	if err != nil {
		return nil, err
	}

	var res PersonProfile
	err = json.Unmarshal(resp, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}
