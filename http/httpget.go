package http

import (
	"encoding/json"
	"time"

	"github.com/Jeffail/gabs"
	"github.com/parnurzeal/gorequest"
)

func UrlPostGetJsonObj(url string, typeStr string, paramMap map[string]interface{}, timeout int64) (*gabs.Container, string, error) {
	param, err := json.Marshal(paramMap)
	if err != nil {
		return nil, "", err
	}
	request := gorequest.New().Timeout(time.Duration(timeout) * time.Second)
	_, body, errs := request.Post(url).Type(typeStr).Send(string(param)).End()
	if errs != nil {
		return nil, body, errs[0]
	}
	jsonParsed, err := gabs.ParseJSON([]byte(body))
	if err != nil {
		return nil, body, err
	}
	return jsonParsed, body, nil
}
