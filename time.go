package utils

import (
	"fmt"
	"strings"
	"time"
)

func GoStdTime() string {
	return "2006-01-02 15:04:05"
}

func GoStdUnixDate() string {
	return "Mon Jan _2 15:04:05 MST 2006"
}

func GetTodayDate() string {
	return time.Now().Format("20060102")
}

func GoStdRubyDate() string {
	return "Mon Jan 02 15:04:05 -0700 2006"
}

func GetTimeHHMM() string {
	return time.Now().Format("15_04")
}

func GetTmStr(tm time.Time, format string) string {
	patterns := []string{
		"y", "2006",
		"m", "01",
		"d", "02",

		"Y", "2006",
		"M", "01",
		"D", "02",

		"h", "03", //12小时制
		"H", "15", //24小时制

		"i", "04",
		"s", "05",

		"t", "pm",
		"T", "PM",
	}
	return convStr(tm, format, patterns)
}

func GetTmShortStr(tm time.Time, format string) string {
	patterns := []string{
		"y", "06",
		"m", "1",
		"d", "2",

		"Y", "06",
		"M", "1",
		"D", "2",

		"h", "3", //12小时制
		"H", "15", //24小时制

		"i", "4",
		"s", "5",

		"t", "pm",
		"T", "PM",
	}
	return convStr(tm, format, patterns)
}

func convStr(tm time.Time, format string, patterns []string) string {
	replacer := strings.NewReplacer(patterns...)
	strfmt := replacer.Replace(format)
	return tm.Format(strfmt)
}

func GenUnixTime() int {
	return int(time.Now().Unix())
}

func GenUnixTime13() int {
	return int(time.Now().UnixNano() / 1000000)
}

func GetLocaltimeStr() string {
	now := time.Now().Local()
	year, mon, day := now.Date()
	hour, min, sec := now.Clock()
	zone, _ := now.Zone()
	return fmt.Sprintf("%d-%d-%d %02d:%02d:%02d %s", year, mon, day, hour, min, sec, zone)
}

func GetGmtimeStr() string {
	now := time.Now()
	year, mon, day := now.UTC().Date()
	hour, min, sec := now.UTC().Clock()
	zone, _ := now.UTC().Zone()
	return fmt.Sprintf("%d-%d-%d %02d:%02d:%02d %s", year, mon, day, hour, min, sec, zone)
}

func GetUnixTimeStr(ut int64, format string) string {
	t := time.Unix(ut, 0)
	return GetTmStr(t, format)
}

func GetUnixTimeShortStr(ut int64, format string) string {
	t := time.Unix(ut, 0)
	return GetTmShortStr(t, format)
}
