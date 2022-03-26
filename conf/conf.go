package conf

import (
	"log"

	"github.com/robfig/config"
)

type Kconf struct {
	configFile string
	c          *config.Config
}

func NewKconf(f string) (*Kconf, error) {
	tmpc, err := config.ReadDefault(f)
	if err != nil {
		log.Println("conf read fail:", err)
		return nil, err
	}

	return &Kconf{
		configFile: f,
		c:          tmpc,
	}, nil
}

func (ts *Kconf) GetString(section string, option string) (value string, err error) {
	return ts.c.String(section, option)
}
func (ts *Kconf) GetInt(section string, option string) (value int, err error) {
	return ts.c.Int(section, option)
}
func (ts *Kconf) GetInt32(section string, option string) (value int32, err error) {
	i, err := ts.GetInt(section, option)
	return int32(i), err
}
func (ts *Kconf) GetInt64(section string, option string) (value int64, err error) {
	i, err := ts.GetInt(section, option)
	return int64(i), err
}
func (ts *Kconf) GetFloat(section string, option string) (value float64, err error) {
	return ts.c.Float(section, option)
}
func (ts *Kconf) GetBool(section string, option string) (value bool, err error) {
	return ts.c.Bool(section, option)
}
