package gomini

import (
	"strings"
	"time"
)

func webTime(t time.Time) string {
	ftime := t.Format(time.RFC1123)
	if strings.HasSuffix(ftime, "UTC") {
		ftime = ftime[0:len(ftime)-3] + "GMT"
	}
	return ftime
}
func GetTimeAgo(t int64) (s string) {
	tt := time.Now().Unix() - t

	if tt < 60 {
		s = GetInt64Str(tt) + " 秒以前"
	} else if tt < 3600 {
		m := tt / 60
		s = GetInt64Str(m) + " 分钟以前"
	} else if tt < 86400 {
		m := tt / 3600
		s = GetInt64Str(m) + " 小时以前"
	} else if tt < 2592000 {
		m := tt / 86400
		s = GetInt64Str(m) + " 天以前"
	} else if tt < 2592000*12 {
		m := tt / 2592000
		s = GetInt64Str(m) + " 月以前"
	} else {
		m := tt / (2592000 * 12)
		s = GetInt64Str(m) + " 年以前"
	}

	return
}
