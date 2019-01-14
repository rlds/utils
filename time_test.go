package utils_test

import (
	"time"
	"testing"
	. "../utils"
)

func TestGoStdTime(t *testing.T) {
	if GoStdTime() != "2006-01-02 15:04:05" {
		t.Error("should be 2006-01-02 15:04:05")
	}
}

func TestGoStdUnixDate(t *testing.T) {
	if GoStdUnixDate() != "Mon Jan _2 15:04:05 MST 2006" {
		t.Error("should be Mon Jan _2 15:04:05 MST 2006")
	}
}

func TestGoStdRubyDate(t *testing.T) {
	if GoStdRubyDate() != "Mon Jan 02 15:04:05 -0700 2006" {
		t.Error("should be Mon Jan 02 15:04:05 -0700 2006")
	}
}

func TestGetTmStr(t *testing.T) {
	str := GetTmStr(time.Now(), "2006-01")
	if str != "2016-05" {
		t.Error("should be 2016-05")
	}
}

func TestGetTmShortStr(t *testing.T) {
	str := GetTmShortStr(time.Now(), "06-1")
	if str != "16-5" {
		t.Error("should be 16-5")
	}
}

func TestGetUnixTimeStr(t *testing.T) {
	if GetUnixTimeStr(1463647489, "2006-01-02") != "2016-05-19" {
		t.Error("should be 2016-05-19")
	}
}

func TestGetUnixTimeShortStr(t *testing.T) {
	if GetUnixTimeStr(1463647489, "06-1-2") != "16-5-19" {
		t.Error("should be 16-5-19")
	}
}
