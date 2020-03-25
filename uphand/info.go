package uphand

import (
	"bufio"
	"encoding/json"
	"fmt"
	"image"
	"net/http"
	"os"

	"github.com/hzde0128/goimg/imghand"
)

// Info 获取图片信息
func Info(w http.ResponseWriter, r *http.Request) {

	// 响应返回
	res := new(UpdateResponse)

	// 获取要图片id
	imgid := r.FormValue("imgid")

	// 获取裁剪后图像的宽度、高度
	width := imghand.StringToInt(r.FormValue("w"))  // 宽度
	height := imghand.StringToInt(r.FormValue("h")) // 高度

	// 组合文件完整路径
	filePath := imghand.URLParse(imgid)
	if filePath == "" {

		res.Code = StatusURLNotFound
		res.Msg = StatusText(StatusURLNotFound)
		w.Write(ResponseJSON(res))

		return
	}

	if width != 0 || height != 0 {
		filePath = fmt.Sprintf("%s_%d_%d", filePath, width, height)
	}

	fimg, err := os.Open(filePath)
	if err != nil {

		res.Code = StatusImgNotFound
		res.Msg = StatusText(StatusImgNotFound)
		w.Write(ResponseJSON(res))

		return
	}
	defer fimg.Close()

	bufimg := bufio.NewReader(fimg)
	_, imgtype, err := image.Decode(bufimg)
	if err != nil {

		res.Code = StatusImgNotFound
		res.Msg = StatusText(StatusImgNotFound)
		w.Write(ResponseJSON(res))

		return
	}

	finfo, _ := fimg.Stat()

	imgstr := imgid + "." + imgtype
	res.Success = true
	res.Code = StatusOK
	res.Msg = StatusText(StatusOK)
	res.Data.Imgid = imgid
	res.Data.Mime = imgtype
	res.Data.Size = finfo.Size()
	res.Data.ImgStr = imgstr
	w.Write(ResponseJSON(res))

}

// StatusCode 状态码
func StatusCode(w http.ResponseWriter, r *http.Request) {
	data, _ := json.Marshal(GetStatusText())
	w.Write(data)
}
