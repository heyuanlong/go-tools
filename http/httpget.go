package http

import (
	//"encoding/json"
	"time"

	"github.com/Jeffail/gabs"
	"github.com/parnurzeal/gorequest"
)

// GoRequest supports
//
//    "text/html" uses "html"
//    "application/json" uses "json"
//    "application/xml" uses "xml"
//    "text/plain" uses "text"
//    "application/x-www-form-urlencoded" uses "urlencoded", "form" or "form-data"
//
// timeout Second
func UrlPostGetJsonObj(url string, typeStr string, paramMap map[string]interface{}, timeout int64) (*gabs.Container, string, error) {
	// param, err := json.Marshal(paramMap)
	// if err != nil {
	// 	return nil, "", err
	// }
	request := gorequest.New().Timeout(time.Duration(timeout) * time.Second)
	_, body, errs := request.Post(url).Type(typeStr).Send(paramMap).End()
	if errs != nil {
		return nil, body, errs[0]
	}
	jsonParsed, err := gabs.ParseJSON([]byte(body))
	if err != nil {
		return nil, body, err
	}
	return jsonParsed, body, nil
}

func UrlGetGetJsonObj(url string, timeout int64) (*gabs.Container, string, error) {
	request := gorequest.New().Timeout(time.Duration(timeout) * time.Second)
	_, body, errs := request.Get(url).End()
	if errs != nil {
		return nil, body, errs[0]
	}
	jsonParsed, err := gabs.ParseJSON([]byte(body))
	if err != nil {
		return nil, body, err
	}
	return jsonParsed, body, nil
}
