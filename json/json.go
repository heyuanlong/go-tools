package json

import (
	"errors"

	"github.com/Jeffail/gabs"
)

func GetJsonObj(str string) (*gabs.Container, error) {
	jsonParsed, err := gabs.ParseJSON([]byte(str))
	if err != nil {
		return nil, err
	}

	return jsonParsed, nil
}

//----------------------------------------------------------------
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

func GabsGetJsonBool(jsonParsed *gabs.Container, path string) (bool, error) {
	v, ok := jsonParsed.Path(path).Data().(bool)
	if !ok {
		return false, errors.New("get value fail")
	}

	return v, nil
}

//----------------------------------------------------------------

func GetJsonString(str string, path string) (string, error) {
	jsonParsed, err := GetJsonObj(str)
	if err != nil {
		return "", err
	}
	v, ok := jsonParsed.Path(path).Data().(string)
	if !ok {
		return "", errors.New("get value fail")
	}

	return v, nil
}

func GetJsonFloat64(str string, path string) (float64, error) {
	jsonParsed, err := GetJsonObj(str)
	if err != nil {
		return 0, err
	}
	v, ok := jsonParsed.Path(path).Data().(float64)
	if !ok {
		return 0, errors.New("get value fail")
	}

	return v, nil
}

func GetJsonInt64(str string, path string) (int64, error) {
	jsonParsed, err := GetJsonObj(str)
	if err != nil {
		return 0, err
	}
	v, ok := jsonParsed.Path(path).Data().(float64)
	if !ok {
		return 0, errors.New("get value fail")
	}

	return int64(v), nil
}

func GetJsonbool(str string, path string) (bool, error) {
	jsonParsed, err := GetJsonObj(str)
	if err != nil {
		return false, err
	}
	v, ok := jsonParsed.Path(path).Data().(bool)
	if !ok {
		return false, errors.New("get value fail")
	}

	return v, nil
}
