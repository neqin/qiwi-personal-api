package qiwi

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	MaxIdleConnections int    = 20000
	RequestTimeout     int    = 25
	ApiUrl             string = "https://edge.qiwi.com"
)

type QiwiPersonalApi struct {
	httpClient *http.Client
	apiKey     string
}

func NewQiwiPersonalApiWithHttpClient(apiKey string, client *http.Client) *QiwiPersonalApi {
	qiwi := new(QiwiPersonalApi)
	qiwi.httpClient = client
	qiwi.apiKey = apiKey
	return qiwi
}

func NewQiwiPersonalApi(apiKey string) *QiwiPersonalApi {
	client := &http.Client{
		Transport: &http.Transport{
			MaxIdleConnsPerHost: MaxIdleConnections,
		},
		Timeout: time.Duration(RequestTimeout) * time.Second,
	}

	return NewQiwiPersonalApiWithHttpClient(apiKey, client)
}

func (qiwi *QiwiPersonalApi) newRequest(apiKey, method, spath string, data map[string]interface{}) (req *http.Request, err error) {
	var path = ApiUrl + spath
	var body io.Reader
	if len(data) > 0 {
		s, err := json.Marshal(data)
		if err != nil {
			return nil, err
		}
		body = bytes.NewBuffer(s)
	}
	req, err = http.NewRequest(method, path, body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	return req, err
}

func (qiwi *QiwiPersonalApi) sendRequest(apiKey, method, spath string, data map[string]interface{}) (body []byte, err error) {
	req, err := qiwi.newRequest(apiKey, method, spath, data)
	response, err := qiwi.httpClient.Do(req)
	if err != nil && response == nil {
		return nil, err
	} else {
		body, err := ioutil.ReadAll(response.Body)
		//log.Println(string(body))
		if err != nil {
			if response.Body != nil {
				response.Body.Close()
			}
			return nil, err
		}
		if response.StatusCode != 200 {
			return nil, errors.New(response.Status)
		}
		if response.Body != nil {
			response.Body.Close()
		}
		return body, nil
	}
}

type Currency int

const (
	RUB Currency = 643
	USD Currency = 840
	EUR Currency = 978
)

type IdentificationLevel string

const (
	SIMPLE    IdentificationLevel = "SIMPLE"
	VERIFIED  IdentificationLevel = "VERIFIED"
	FULL      IdentificationLevel = "FULL"
	ANONYMOUS IdentificationLevel = "ANONYMOUS"
)
