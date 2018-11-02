package e

import (
	"fmt"
	"strconv"
	"time"
)

const (
	BASE_FORMAT_D = "2006-01-02"
	BASE_FORMAT_S = "2006-01-02 15:04:05"
	BASE_FORMAT_MIN  = 60
	BASE_FORMAT_HOUR = 60 * 60
	BASE_FORMAT_CST  = 60 * 60 * 8
	BASE_FORMAT_DAY  = 60 * 60 * 24
	BASE_FORMAT_WEEK = 60 * 60 * 24 * 7
	BASE_FORMAT_MONTH = 60 * 60 * 24 * 30
	BASE_FORMAT_YEAR  = 60 * 60 * 24 * 365
	BASE_FORMAT_TEN_YEAR = 60 * 60 * 24 * 365 * 10
)

func StrToTime(t string) (unixTime int64) {
	timeTemp, err := time.Parse(BASE_FORMAT_S, t)
	if err == nil {
		unixTime = timeTemp.Unix() - BASE_FORMAT_CST
	}
	return
}

func StrTimeToWeek(t string) (result string) {

	var timeStr string
	timeStr = t+" 00:00:00"
	week := time.Unix(StrToTime(timeStr), 0).Weekday().String()
	switch week {
	case "Sunday":
		result = t + " 周日"
	case "Monday":
		result = t + " 周一"
	case "Tuesday":
		result = t + " 周二"
	case "Wednesday":
		result = t + " 周三"
	case "Thursday":
		result = t + " 周四"
	case "Friday":
		result = t + " 周五"
	case "Saturday":
		result = t + " 周六"
	default:
		result = t
	}

	return
}

func TimeToStr(t int64) (strTime string) {
	strTime = time.Unix(t, 0).Format(BASE_FORMAT_S)
	return
}

func TimeFormatShow(t int64, tail string) (strTime string) {

	nowTime := time.Now().Unix()
	diff := nowTime - t
	if diff < BASE_FORMAT_HOUR { // 分钟 < 60 * 60
		strTime = strconv.FormatInt(diff / BASE_FORMAT_MIN, 10) + "分钟" + tail
	}else if diff < BASE_FORMAT_DAY { // 小时 < 60 * 60 * 24
		strTime = strconv.FormatInt(diff / BASE_FORMAT_HOUR, 10) + "小时" + tail
	}else if diff < BASE_FORMAT_WEEK { // 天   < 60 * 60 * 24 * 7
		strTime = strconv.FormatInt(diff / BASE_FORMAT_DAY, 10) + "天" + tail
	}else if diff < BASE_FORMAT_MONTH { // 周   < 60 * 60 * 24 * 30
		strTime = strconv.FormatInt(diff / BASE_FORMAT_WEEK, 10) + "周" + tail
	}else if diff < BASE_FORMAT_YEAR { // 月   < 60 * 60 * 24 * 365
		strTime = strconv.FormatInt(diff / BASE_FORMAT_MONTH, 10) + "月" + tail
	}else if diff < BASE_FORMAT_TEN_YEAR { // 年   < 60 * 60 * 24 * 365 * 10
		strTime = strconv.FormatInt(diff / BASE_FORMAT_YEAR, 10) + "年" + tail
	}else { // 盘古开荒前
		strTime = "盘古开荒" + tail
	}

	return
}


func testTimeFunc() {

	// Time格式输出
	nt := time.Now()
	fmt.Println(nt.Format(BASE_FORMAT_S))
	// 2018-10-19 15:35:21.4236437 +0800 CST m=+0.004974601

	// 时间戳输出
	timestamp := time.Now().Unix()
	fmt.Println(timestamp)
	// 1539934521

	// 时间戳转换 时间格式
	a := TimeToStr(timestamp)
	fmt.Printf("timeToStr: from [%d] to [%s] \n", timestamp, a)
	// timeToStr: from [1539934521] to [2018-10-19 15:35:21]

	// 时间格式 转 时间戳
	b := StrToTime(a)
	fmt.Printf("strToTime: from [%s] to [%d] \n", a, b)
	// strToTime: from [2018-10-19 15:35:21] to [1539963321]

	// 时间格式化
	var t int64
	tail := "之前"
	t = 14730350
	fmt.Println(TimeFormatShow(t, tail))
}