package common

import (
	"time"
)

//秒
func GetTimes() int64 {
	return time.Now().Unix()
}

//这方式比较特别，按照123456来记忆吧：01月02号 下午3点04分05秒 2006年
func GetTimesString() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

//纳秒
func GetTimesNano() int64 {
	return time.Now().UnixNano()
}

//日期转时间戳
func DateToTimestamp(str string) (int64, error) {
	loc, err := time.LoadLocation("Local")
	if err != nil {
		return 0, err
	}
	tm, err := time.ParseInLocation("2006-01-02", str, loc)
	if err != nil {
		return 0, err
	}
	return tm.Unix(), nil
}

//时间转时间戳
func DateTimeToTimestamp(str string) (int64, error) {
	loc, err := time.LoadLocation("Local")
	if err != nil {
		return 0, err
	}
	tm, err := time.ParseInLocation("2006-01-02 15:04:05", str, loc)
	if err != nil {
		return 0, err
	}
	return tm.Unix(), nil
}

//时间戳转日期
func TimestampToDate(tm int64) string {
	return time.Unix(tm, 0).Format("2006-01-02")
}

//时间戳转时间
func TimestampToDateTime(tm int64) string {
	return time.Unix(tm, 0).Format("2006-01-02 15:04:05")
}

func SqlDateTimeToTimestamp(str string) (int64, error) {
	loc, err := time.LoadLocation("Local")
	if err != nil {
		return 0, err
	}
	tm, err := time.ParseInLocation("2006-01-02T15:04:05+08:00", str, loc)
	if err != nil {
		return 0, err
	}
	return tm.Unix(), nil
}

func SqlDateTimeToDate(str string) string {
	loc, err := time.LoadLocation("Local")
	if err != nil {
		return "2006-01-02"
	}
	tm, err := time.ParseInLocation("2006-01-02T15:04:05+08:00", str, loc)
	if err != nil {
		return "2006-01-02"
	}
	return tm.Format("2006-01-02")
}

func SqlDateTimeToDateTime(str string) string {
	loc, err := time.LoadLocation("Local")
	if err != nil {
		return "2006-01-02 15:04:05"
	}
	tm, err := time.ParseInLocation("2006-01-02T15:04:05+08:00", str, loc)
	if err != nil {
		return "2006-01-02 15:04:05"
	}
	return tm.Format("2006-01-02 15:04:05")
}

//timestamp=4*3600,表示获取当前时间到下次凌晨4点的秒数
func GetTimerNext(timestamp int64) int64 {
	t := time.Now()
	nowTime := t.Unix()
	y, m, d := t.Date()
	beginDayTime := time.Date(y, m, d, 0, 0, 0, 0, t.Location()).Unix()

	var timer int64
	if (beginDayTime + timestamp) > nowTime {
		timer = (beginDayTime + timestamp) - nowTime
	} else {
		timer = (beginDayTime + timestamp + 3600*24) - nowTime
	}
	return timer
}
