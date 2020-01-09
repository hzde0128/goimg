package imghand

import (

	"log"
	"net/url"
	"strconv"
	"strings"

	"github.com/hzde0128/goimg/config"
)

// 匹配是否是 md5 的长度
func IsMD5Path(str string) bool {
	return regexpUrlParse.MatchString(str)
}

// 路径部分排序做目录
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

// 组合文件目录路径
func JoinPath(md5_str string) string {

	// 路径部分排序做目录
	sortPath := SortPath([]byte(md5_str[:5]))

	var str = strings.Builder{}

	str.WriteString(config.PathImg())
	str.WriteString(sortPath)
	str.WriteString("/")
	str.WriteString(md5_str[0:32])

	// 配置文件目录/短目录/md5/md5图片
	return str.String()

}

// 进行 url 部分解析 - md5，并组合文件完整路径
func UrlParse(md5_url string) string {

	if md5_url == "" {
		return ""
	}

	// 进行 url 解析
	parse, err := url.Parse(md5_url)
	if err != nil {
		return ""
	}

	parsePath := parse.Path

	if len(parsePath) < 32 {
		return ""
	}

	// 匹配是否是 md5 的长度
	if !IsMD5Path(parsePath) {
		log.Printf("%s 图片不匹配", parsePath)
		return ""
	}

	// 组合文件完整路径
	return JoinPath(parsePath) + "/" + parsePath

}

// 字符串的数字转int
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
