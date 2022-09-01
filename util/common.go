package util

import (
	"github.com/spf13/cast"
	"time"
)

func GetCurrentlyTime() string {
	return cast.ToString(time.Now().Format("2006-01-02 15:04:05"))
}
func GetTogetherDays() int {
	beginDate := cast.ToTimeInDefaultLocation("2022-06-11", time.Local)
	nowDate := cast.ToTime(time.Now())
	return getTimesSubDay(nowDate, beginDate)
}

func getTimesSubDay(time1, time2 time.Time) int {
	sub := time1.Sub(time2)
	days := sub.Hours() / 24
	return int(days)
}
