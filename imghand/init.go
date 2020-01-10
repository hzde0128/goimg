package imghand

// 图片处理

import (
	"log"
	"regexp"
	"image"
)

var regexpUrlParse *regexp.Regexp

var noImg *image.RGBA

func init() {

	var err error
	// 初始化正则表达式 图片正则，32位md5+时间戳+后缀
	regexpUrlParse, err = regexp.Compile(`[a-z0-9]{32}\_\d+\.(jpeg|jpg|png|gif)`)
	if err != nil {
		log.Fatalln("regexpUrlParse:", err)
	}

	// 创建 RGBA 画板大小 - 用于找不到图片时用
	noImg = image.NewRGBA(image.Rect(0, 0, 400, 400))

}
