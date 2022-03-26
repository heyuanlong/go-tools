package common

import (
	"bytes"
	"errors"
	"fmt"
	"math/rand"
	"net/url"
	"sort"
	"strconv"
	"strings"
)

//0：数字+大小写字母，1：数字+小写字母，2：数字+大写字母，3：数字，4：小写字母
func GetRandomString(lens int, types int) string {
	var str string
	if types == 1 {
		str = "0123456789abcdefghijklmnopqrstuvwxyz"
	} else if types == 2 {
		str = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	} else if types == 3 {
		str = "0123456789"
	} else if types == 4 {
		str = "abcdefghijklmnopqrstuvwxyz"
	} else {
		str = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	}

	bytes := []byte(str)
	result := []byte{}
	for i := 0; i < lens; i++ {
		result = append(result, bytes[rand.Intn(len(bytes))])
	}
	return string(result)
}

//map转为url参数，带升序排序
//seq分隔字符
func ChangeMapToURLParam(param map[string]interface{}, sep string) (string, error) {
	paramM := make(map[string]string)
	for k, v := range param {
		switch val := v.(type) {
		case string:
			paramM[k] = val
		case bool:
			paramM[k] = strconv.FormatBool(val)
		case int:
			paramM[k] = strconv.FormatInt(int64(val), 10)
		case int32:
			paramM[k] = strconv.FormatInt(int64(val), 10)
		case int64:
			paramM[k] = strconv.FormatInt(int64(val), 10)
		case float32:
			paramM[k] = strconv.FormatFloat(float64(val), 'f', -1, 64)
		case float64:
			paramM[k] = strconv.FormatFloat(float64(val), 'f', -1, 64)
		default:
			//klog.Warn.Println(k, v)
			return "", errors.New("not find value type")
		}
	}
	lens := len(paramM)
	keys := make([]string, 0, lens)
	for k := range paramM {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	var b bytes.Buffer

	for i, k := range keys {
		if i == (lens - 1) {
			fmt.Fprintf(&b, "%s=%s", k, paramM[k])
		} else {
			fmt.Fprintf(&b, "%s=%s%s", k, paramM[k], sep)
		}
	}
	return b.String(), nil
}

func BuildUrlParam(sUrl string, params map[string]interface{}) string {
	var buf strings.Builder
	buf.WriteString(sUrl)
	buf.WriteByte('?')

	for k, v := range params {
		buf.WriteString(url.QueryEscape(k))
		buf.WriteByte('=')
		buf.WriteString(url.QueryEscape(fmt.Sprint(v)))
		buf.WriteByte('&')
	}

	return buf.String()
}
