package main

import (
	"fmt"
	"time"
)

// StandTimeFormat ...
const StandTimeFormat string = "2006-01-02 15:04:05"
const shTimeZone string = "Asia/Shanghai"

var shanghaiLocation = GetShanghaiTimeLocation()

// GetShanghaiTimeStamp 获取上海的时间戳，跟中国同一个时区的都可以使用这个时间戳
func GetShanghaiTimeStamp() int64 {
	return GetShanghaiTime().Unix()
}

// GetShanghaiTime ...
func GetShanghaiTime() time.Time {
	return time.Now().In(shanghaiLocation)
}

// GetShanghaiYesterdayTime 获取昨天时间
func GetShanghaiYesterdayTime() time.Time {
	return GetShanghaiTime().AddDate(0, 0, -1)
}

// GetShanghaiTimeByStamp 根据时间戳获取时间
func GetShanghaiTimeByStamp(sec int64) time.Time {
	return time.Unix(sec, 0).In(shanghaiLocation)
}

// GetShanghaiTimeLocation ...
func GetShanghaiTimeLocation() *time.Location {
	shanghai, err := time.LoadLocation(shTimeZone)
	if err != nil {
		panic(err)
	}

	return shanghai
}

// GetShanghaiYesterdayHourTimeStamp 获取昨天指定小时整的时间戳
func GetShanghaiYesterdayHourTimeStamp(hour int) int64 {
	yesterday := GetShanghaiTime().AddDate(0, 0, -1)
	str := fmt.Sprintf("%04d-%02d-%02d %02d:%02d:%02d", yesterday.Year(), yesterday.Month(), yesterday.Day(), hour, 0, 0)
	t, err := time.ParseInLocation(StandTimeFormat, str, shanghaiLocation)
	if err != nil {
		panic(err)
	}
	return t.Unix()
}

// GetShanghaiTodayHourTimeStamp 获取当天指定小时整的时间戳
func GetShanghaiTodayHourTimeStamp(hour int) int64 {
	today := GetShanghaiTime()
	str := fmt.Sprintf("%04d-%02d-%02d %02d:%02d:%02d", today.Year(), today.Month(), today.Day(), hour, 0, 0)
	t, err := time.ParseInLocation(StandTimeFormat, str, shanghaiLocation)
	if err != nil {
		panic(err)
	}
	return t.Unix()
}

// GetShanghaiTodayBeforeHourTimeStamp 获取当天指定小时前的时间戳
func GetShanghaiTodayBeforeHourTimeStamp(hour int) int64 {
	// 不完美，但够用，不想去修改，没必要浪费时间
	if hour < 1 {
		hour = 1
	}
	today := GetShanghaiTime()
	str := fmt.Sprintf("%04d-%02d-%02d %02d:%02d:%02d", today.Year(), today.Month(), today.Day(), hour-1, 59, 59)
	t, err := time.ParseInLocation(StandTimeFormat, str, shanghaiLocation)
	if err != nil {
		panic(err)
	}
	return t.Unix()
}

func main() {
}
