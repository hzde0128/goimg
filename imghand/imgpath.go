package imghand

import (
	"log"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/hzde0128/goimg/config"
)

// IsMD5Path 匹配是否是 md5 的长度
func IsMD5Path(str string) bool {
	return regexpURLParse.MatchString(str)
}

// SortPath 路径部分排序做目录
func SortPath(str []byte) string {

	// 对 byte 进行排序
	strLen := len(str)
	for i := 0; i < strLen; i++ {
		for j := 1 + i; j < strLen; j++ {
			if str[i] > str[j] {
				str[i], str[j] = str[j], str[i]
			}
		}
	}

	// 对 byte 依次组成数字符串
	var ret = strings.Builder{}

	for i := 0; i < strLen; i++ {
		ret.WriteString(strconv.Itoa(int(str[i])))
	}

	return ret.String()
}

// JoinPath 组合文件目录路径
func JoinPath(md5Str string) string {

	// 路径部分排序做目录
	sortPath := SortPath([]byte(md5Str[:5]))

	var str = strings.Builder{}

	str.WriteString(config.PathImg())
	str.WriteString(sortPath)
	str.WriteString("/")
	str.WriteString(md5Str[0:32])

	// 配置文件目录/短目录/md5/md5图片
	return str.String()

}

// JoinPath1 根据时间戳计算目录
func JoinPath1(timeStamp string) string {

	timestamp, err := strconv.Atoi(timeStamp)
	if err != nil {
		log.Printf("时间戳转换整型失败:%v", err)
	}
	dateTime := time.Unix(int64(timestamp), 0).Format("20060102")
	var str = strings.Builder{}

	str.WriteString(config.PathImg())
	str.WriteString("/")
	str.WriteString(dateTime)
	return str.String()
}

// URLParse 进行 url 部分解析 - md5，并组合文件完整路径
func URLParse(md5Url string) string {

	if md5Url == "" {
		return ""
	}

	// 进行 url 解析
	parse, err := url.Parse(md5Url)

	if err != nil {
		return ""
	}

	parsePath := parse.Path

	// 根据时间戳获取对应的目录
	// fb22e5e61756acd6c070065139186b7c_1578646180.jpg
	startIndex := strings.LastIndex(md5Url, "_")
	endIndex := strings.LastIndex(md5Url, ".")
	timeStamp := md5Url[startIndex+1 : endIndex]
	// timeStampInt, _ := strconv.Atoi(timeStamp)

	// timestamp转日期
	// dateTime := time.Unix(int64(timeStampInt),0).Format("20060102")
	// log.Printf("dateTime:%v",dateTime)
	if len(parsePath) < 32 {
		return ""
	}

	// 匹配是否是 md5 的长度
	if !IsMD5Path(parsePath) {
		log.Printf("%s 图片不匹配", parsePath)
		return ""
	}

	// 组合文件完整路径
	// log.Printf("parsePath:%v\n",parsePath)
	return JoinPath1(timeStamp) + "/" + parsePath

}

// StringToInt 字符串的数字转int
func StringToInt(str string) int {
	if str == "" {
		return 0
	}

	toint, err := strconv.Atoi(str)
	if err != nil {
		return 0
	}

	if toint < 0 {
		return 0
	}

	return toint
}
