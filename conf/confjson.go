package conf

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/Jeffail/gabs"
)

type JsonConf struct {
	configFile string
	Container  *gabs.Container
}

func NewJsonConf(f string) (*JsonConf, error) {
	obj := &JsonConf{
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

func (ts *JsonConf) GetString(path string) (value string, err error) {
	v, ok := ts.Container.Path(path).Data().(string)
	if !ok {
		return "", errors.New("get value fail")
	}
	return v, nil
}
func (ts *JsonConf) GetFloat64(path string) (value float64, err error) {
	v, ok := ts.Container.Path(path).Data().(float64)
	if !ok {
		return 0, errors.New("get value fail")
	}
	return v, nil
}
func (ts *JsonConf) GetInt64(path string) (value int64, err error) {
	v, ok := ts.Container.Path(path).Data().(float64)
	if !ok {
		return 0, errors.New("get value fail")
	}
	return int64(v), nil
}

func (ts *JsonConf) GetInt(path string) (value int, err error) {
	v, ok := ts.Container.Path(path).Data().(float64)
	if !ok {
		return 0, errors.New("get value fail")
	}
	return int(v), nil
}

func (ts *JsonConf) GetBool(path string) (value bool, err error) {
	v, ok := ts.Container.Path(path).Data().(bool)
	if !ok {
		return false, errors.New("get value fail")
	}
	return v, nil
}

func (ts *JsonConf) GetMap(path string) (map[string]string, error) {
	result := make(map[string]string)
	cMap, err := ts.Container.Path(path).ChildrenMap()
	if err != nil {
		return result, err
	}
	for key, child := range cMap {
		result[key] = ts.ConvertTypeToString(child.Data())
	}
	return result, nil
}

func (ts *JsonConf) GetArrayMap(path string, slave string, rmPath ...string) (result []map[string]map[string]string, err error) {
	cMap, err := ts.Container.Path(path).ChildrenMap()
	if err != nil {
		return
	}
	rmMap := make(map[string]bool)
	for _, val := range rmPath {
		rmMap[val] = true
	}
	//遍历判断是否有除了数据库连接参数外的其他数据库
	for key, _ := range cMap {
		//默认数据库的读库
		if key == slave {
			res, mapErr := ts.GetArrayToMap(fmt.Sprintf("%s.%s", path, key), true, slave, rmPath...)
			if mapErr != nil {
				err = mapErr
				return
			}
			result = append(result, res)
		} else {
			//其他数据库
			if _, ok := rmMap[key]; !ok {
				res, mapErr := ts.GetArrayToMap(fmt.Sprintf("%s.%s", path, key), false, slave, rmPath...)
				if mapErr != nil {
					err = mapErr
					return
				}
				result = append(result, res)
			}
		}
	}
	return
}

func (ts *JsonConf) GetArrayToMap(path string, isRead bool, slave string, str ...string) (map[string]map[string]string, error) {

	result := make(map[string]map[string]string)
	if isRead {
		//读写分离下获取读库的各种连接参数
		readChild, _ := ts.Container.Path(path).Children()
		for readChildIndex, readChildVal := range readChild {
			readChildIndexKey := fmt.Sprintf("default_read_0_%d", readChildIndex)
			for _, key := range str {
				if !readChildVal.Exists(key) {
					continue
				}

				tmpVal := strings.TrimSpace(ts.ConvertTypeToString(readChildVal.Path(key).Data()))
				if readChildIndexVal, ok := result[readChildIndexKey]; ok {
					readChildIndexVal[key] = tmpVal
				} else {
					result[readChildIndexKey] = map[string]string{key: tmpVal}
				}

			}
		}
	} else {
		//读写分离下获取写库同名下的后缀数字
		childMap, _ := ts.Container.Path(path).ChildrenMap()
		tmpName := strings.Split(path, ".")
		nameIndex := strings.Split(tmpName[len(tmpName)-1], "_")
		countIndex := len(nameIndex)
		var childMapIndex int64 = 0
		if countIndex > 1 {
			tmpIndex, err := strconv.ParseInt(nameIndex[countIndex-1], 10, 64)
			if err != nil {
				return result, err
			}
			childMapIndex = tmpIndex
		}
		//遍历写库所有参数
		for childKey, child := range childMap {
			//写库中的读库参数
			if childKey == slave {
				arrChild, _ := ts.Container.Path(fmt.Sprintf("%s.read", path)).Children()
				for arrIndex, arrVal := range arrChild {
					arrIndexKey := fmt.Sprintf("%s_read_%d_%d", nameIndex[0], childMapIndex, arrIndex)
					//判断是否有默认的数据库连接参数
					for _, key := range str {
						//没有默认的数据库连接参数时获取写库的参数
						if !arrVal.Exists(key) {
							//写库也没有默认的数据库连接参数时获取默认数据库的连接参数
							if childData, ok := childMap[key]; ok {
								tmpVal := strings.TrimSpace(ts.ConvertTypeToString(childData.Data()))
								if arrIndexKeyVal, ok := result[arrIndexKey]; ok {
									arrIndexKeyVal[key] = tmpVal
								} else {
									result[arrIndexKey] = map[string]string{key: tmpVal}
								}
							}
							continue
						}

						tmpVal := strings.TrimSpace(ts.ConvertTypeToString(arrVal.Path(key).Data()))
						if arrIndexKeyVal, ok := result[arrIndexKey]; ok {
							arrIndexKeyVal[key] = tmpVal
						} else {
							result[arrIndexKey] = map[string]string{key: tmpVal}
						}
					}
				}
			} else {
				childMapIndexKey := fmt.Sprintf("%s_write_%d_0", nameIndex[0], childMapIndex)
				tmpVal := strings.TrimSpace(ts.ConvertTypeToString(child.Data()))
				if childMapIndexVal, ok := result[childMapIndexKey]; ok {
					childMapIndexVal[childKey] = tmpVal
				} else {
					result[childMapIndexKey] = map[string]string{childKey: tmpVal}
				}
			}
		}
	}
	return result, nil
}

func (ts *JsonConf) ConvertTypeToString(value interface{}) string {
	result := ""
	switch value.(type) {
	case float64:
		result = fmt.Sprint(value.(float64))
	case bool:
		result = strconv.FormatBool(value.(bool))
	case string:
		result = value.(string)
	}
	return result
}
