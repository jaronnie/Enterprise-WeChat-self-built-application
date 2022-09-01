package util

import "testing"

func TestGetCurrentlyTime(t *testing.T) {
	t.Log(GetCurrentlyTime())
}

func TestGetTogetherDays(t *testing.T) {
	days := GetTogetherDays()
	t.Log(days)
}
