package conf

import (
	"errors"
	"log"

	"github.com/Jeffail/gabs"
)

type Kconf struct {
	configFile string
	Container  *gabs.Container
}

func NewKconf(f string) (*Kconf, error) {
	obj := &Kconf{
		configFile: f,
		Container:  nil,
	}
	c, err := gabs.ParseJSONFile(obj.configFile)
	if err != nil {
		log.Println("conf parse fail:", err)
		return nil, err
	}
	obj.Container = c
	return obj, nil
}

func (ts *Kconf) GetString(path string) (value string, err error) {
	v, ok := ts.Container.Path(path).Data().(string)
	if !ok {
		return "", errors.New("get value fail")
	}
	return v, nil
}
func (ts *Kconf) GetFloat64(path string) (value float64, err error) {
	v, ok := ts.Container.Path(path).Data().(float64)
	if !ok {
		return 0, errors.New("get value fail")
	}
	return v, nil
}
func (ts *Kconf) GetInt64(path string) (value int64, err error) {
	v, ok := ts.Container.Path(path).Data().(float64)
	if !ok {
		return 0, errors.New("get value fail")
	}
	return int64(v), nil
}

func (ts *Kconf) GetInt(path string) (value int, err error) {
	v, ok := ts.Container.Path(path).Data().(float64)
	if !ok {
		return 0, errors.New("get value fail")
	}
	return int(v), nil
}

func (ts *Kconf) GetBool(path string) (value bool, err error) {
	v, ok := ts.Container.Path(path).Data().(bool)
	if !ok {
		return false, errors.New("get value fail")
	}
	return v, nil
}
