package http

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/Jeffail/gabs"
	"github.com/parnurzeal/gorequest"
)

func UrlPostGetJsonObj(url string, typeStr string, paramMap map[string]interface{}) (*gabs.Container, string, error) {
	param, err := json.Marshal(paramMap)
	if err != nil {
		return nil, "", err
	}
	request := gorequest.New().Timeout(10 * time.Second)
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

func GabsGetJsonString(jsonParsed *gabs.Container, path string) (string, error) {
	v, ok := jsonParsed.Path(path).Data().(string)
	if !ok {
		return "", errors.New("get value fail")
	}

	return v, nil
}

func GabsGetJsonFloat64(jsonParsed *gabs.Container, path string) (float64, error) {
	v, ok := jsonParsed.Path(path).Data().(float64)
	if !ok {
		return 0, errors.New("get value fail")
	}

	return v, nil
}

func GabsGetJsonInt64(jsonParsed *gabs.Container, path string) (int64, error) {
	v, ok := jsonParsed.Path(path).Data().(float64)
	if !ok {
		return 0, errors.New("get value fail")
	}

	return int64(v), nil
}
