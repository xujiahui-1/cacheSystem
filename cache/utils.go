package cache

import (
	"log"
	"regexp"
	"strconv"
	"strings"
)

const (
	B = 1 << (iota * 10)
	KB
	MB
	GB
	TB
	PB
)

// ParseSize 处理memory size
func ParseSize(size string) (int64, string) {
	//这里的考虑是，当用户输入不规范时，我们1要给出提示，2要给默认值

	//获取单位部分
	re, _ := regexp.Compile("[0-9]+")
	unit := string(re.ReplaceAll([]byte(size), []byte("")))
	//获取数字部分
	number, _ := strconv.ParseInt(strings.Replace(size, unit, "", 1), 10, 64)
	unit = strings.ToUpper(unit)

	var byteNum int64 = 0
	switch unit {
	case "B":
		byteNum = number
	case "KB":
		byteNum = number * KB
	case "MB":
		byteNum = number * MB
	case "GB":
		byteNum = number * GB
	case "TB":
		byteNum = number * TB
	case "PB":
		byteNum = number * PB
	default:
		number = 0

	}
	if number == 0 {
		log.Println("Parse Size 仅支持 B,KB,MB,GB,TB,PB")
		number = 100
		byteNum = number * MB
		unit = "MB"

	}

	return byteNum, strconv.FormatInt(number, 10) + unit
}
