package e

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
	"zimuzu_web_api/pkg/gredis"
)

/**
	截取字符串
 */
func SubString(str string,begin,length int) (substr string) {
	// 将字符串的转换成[]rune
	rs := []rune(str)
	lth := len(rs)

	// 简单的越界判断
	if begin < 0 {
		begin = 0
	}
	if begin >= lth {
		begin = lth
	}
	end := begin + length
	if end > lth {
		end = lth
	}

	if begin == 0 && end < lth {
		substr = string(rs[begin:end]) + "..."
	}else {
		substr = string(rs[begin:end])
	}
	return
}

/**
	补全图片路径
 */
func GetImagPath(path string, typeS string) (ImagPath string) {

	if len(path) > 0 {
		ImagPath = IMG_BASE_URL + "ftp/" + typeS + "_" + path
	}else {
		ImagPath = IMG_BASE_URL + "pic/reso.gif"
	}
	return
}

/**
	trim 掉多余的空格和换行符
 */
func Trim(str string) (newStr string) {

	// 将字符串的转换成[]rune
	strList := []rune(str)
	lth := len(strList)
	star := 0
	end  := lth - 1
	for i:=0; i<lth; i++ {
		if star == i {
			if string(strList[i:i+1]) == " " {
				star ++
			}
		}else {
			if string(strList[i:i+1]) == " " {
				end = i
			}
		}
	}

	if star < end {
		newStr = string(strList[star:end])
	}
	return
}

/**
	默认从Redis取数据 没有取到则从 回调函数取
 */
func RedisGet(redisKey string, id uint, do func(id uint) ) {

	existsTag := gredis.Exists(redisKey)
	if !existsTag {
		do(id)
	}
}

/**
	MD5 字符串
 */
func Md5Str(text string) string {
	ctx := md5.New()
	ctx.Write([]byte(text))
	return hex.EncodeToString(ctx.Sum(nil))
}

/**
	四舍五入 精度保留几位小数 Float64
 */
func NumberFloat64Format(value float64, t int) (result float64) {

	format := "%."+strconv.Itoa(t)+"f"
	result, _ = strconv.ParseFloat(fmt.Sprintf(format, value), 64)
	return
}

func NumberFormat(a int, b int) (result float64) {

	aStr, err := strconv.ParseFloat(strconv.Itoa(a), 64)
	if err != nil {
		aStr = 0
	}
	bStr, err := strconv.ParseFloat(strconv.Itoa(b), 64)
	if err != nil {
		bStr = 0
	}

	return NumberFloat64Format(aStr / bStr, 2)
}
