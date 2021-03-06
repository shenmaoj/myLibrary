/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2019-06-28 11:20 
# @File : time.go
# @Description : 
*/
package utils

import (
	"time"
)

const (
	BASE_TIME_FORMAT_TILL_SEC = "2006-01-02 03:04:05"
	BASE_TIME_STRING_FORMAT   = "2006-01-02 15:04:05"
)

func FormatTime2StringByTemplate(template string, t int64) string {
	return time.Unix(t, 0).Format(template)
}

func GetBeiJingTimeZone() *time.Location {
	var cstZone = time.FixedZone("CST", 8*3600) // 东八
	return cstZone
}

func TimeConvString2Int64(str string) time.Time {
	parse, _ := time.Parse(BASE_TIME_STRING_FORMAT, str)
	// location, _ := time.ParseInLocation(BASE_TIME_STRING_FORMAT, str, time.Local)
	return parse
}

func Int64ConvT2TimeStrTilSec(timeStamp int64) string {
	return time.Unix(timeStamp, 0).Format(BASE_TIME_FORMAT_TILL_SEC)
}
