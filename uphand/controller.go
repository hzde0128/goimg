package uphand

import (
	"bufio"
	"crypto/md5"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/hzde0128/goimg/imghand"
)

// Controller 结构体
type Controller struct {
}

func (c Controller) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.String() == "/favicon.ico" {
		return
	}

	if r.Method == "GET" {
		c.Get(w, r)
		return
	}

	if r.Method == "POST" {
		c.Post(w, r)
		return
	}
}

// Get 显示图片
func (c Controller) Get(w http.ResponseWriter, r *http.Request) {

	urlParse := r.URL.String()

	// 组合文件完整路径

	imgStr := imghand.URLParse(urlParse[1:])

	// 请求/目录
	if imgStr == "" {
		w.Write(showMain())
		return
	}

	// 获取 要裁剪图像的宽度、高度
	width := imghand.StringToInt(r.FormValue("w"))  // 宽度
	height := imghand.StringToInt(r.FormValue("h")) // 高度

	// 加载图片
	imghand.CutImage(w, imgStr, width, height)

}

// Post 上传图片
func (c Controller) Post(w http.ResponseWriter, r *http.Request) {

	// 响应返回
	res := new(UpdateResponse)

	// 上传表单 --------------------------------------
	// 缓冲的大小 - 8M
	r.ParseMultipartForm(1024 << 13)
	// 是上传表单域的名字fileHeader
	upfile, upFileInfo, err := r.FormFile("file")
	if err != nil {
		res.Code = StatusForm
		res.Msg = StatusText(StatusForm)
		w.Write(ResponseJSON(res))
		return
	}
	defer upfile.Close()

	// 图片解码 --------------------------------------
	// 读入缓存
	bufUpFile := bufio.NewReader(upfile)
	// 进行图片的解码
	img, imgtype, err := image.Decode(bufUpFile)
	if err != nil {
		res.Code = StatusImgDecode
		res.Msg = StatusText(StatusImgDecode)
		w.Write(ResponseJSON(res))
		return
	}

	// 判断是否有这个图片类型
	if !imghand.IsType(imgtype) {
		res.Code = StatusImgIsType
		res.Msg = StatusText(StatusImgIsType)
		w.Write(ResponseJSON(res))
		return
	}

	// 设置文件读写下标 --------------------------------
	// 设置下次读写位置（移动文件指针位置）
	_, err = upfile.Seek(0, 0)
	if err != nil {
		res.Code = StatusFileSeek
		res.Msg = StatusText(StatusFileSeek)
		w.Write(ResponseJSON(res))
		return
	}

	// 计算文件的 MD5 值 -----------------------------
	// 初始化 MD5 实例
	md5Hash := md5.New()
	// 读入缓存
	bufFile := bufio.NewReader(upfile)
	_, err = io.Copy(md5Hash, bufFile)
	if err != nil {
		res.Code = StatusFileMd5
		res.Msg = StatusText(StatusFileMd5)
		w.Write(ResponseJSON(res))
		return
	}
	// 进行 MD5 算计，返回 16进制的 byte 数组
	fileMd5FX := md5Hash.Sum(nil)
	fileMd5 := fmt.Sprintf("%x", fileMd5FX)

	// 时间戳,单位为秒
	timeStamp := strconv.Itoa(int(time.Now().Unix()))

	// 目录计算 --------------------------------------
	// 组合文件完整路径
	dirPath := imghand.JoinPath(timeStamp) + "/" // 目录

	// 修改jpeg的后缀为jpg
	if imgtype == "jpeg" {
		imgtype = "jpg"
	}

	// 完整文件路径
	filePath := dirPath + fileMd5 + "_" + timeStamp + "." + imgtype

	// 获取目录信息，并创建目录
	dirInfo, err := os.Stat(dirPath)
	if err != nil {
		err = os.MkdirAll(dirPath, 0755)
		if err != nil {

			res.Code = StatusMkdir
			res.Msg = StatusText(StatusMkdir)
			w.Write(ResponseJSON(res))

			return
		}
	} else {
		if !dirInfo.IsDir() {
			err = os.MkdirAll(dirPath, 0755)
			if err != nil {
				res.Code = StatusMkdir
				res.Msg = StatusText(StatusMkdir)
				w.Write(ResponseJSON(res))
				return
			}
		}
	}

	// 存入文件 --------------------------------------

	_, err = os.Stat(filePath)
	if err != nil {

		// 打开一个文件,文件不存在就会创建
		file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			res.Code = StatusOpenFile
			res.Msg = StatusText(StatusOpenFile)
			w.Write(ResponseJSON(res))
			return
		}

		defer file.Close()

		if imgtype == imghand.PNG {
			err = png.Encode(file, img)
		} else if imgtype == imghand.JPG || imgtype == imghand.JPEG {
			err = jpeg.Encode(file, img, nil)
		} else if imgtype == imghand.GIF {
			// 设置下次读写位置（移动文件指针位置）
			_, err = upfile.Seek(0, 0)
			if err != nil {
				res.Code = StatusFileSeek
				res.Msg = StatusText(StatusFileSeek)
				w.Write(ResponseJSON(res))
				return
			}
			gifimg, giferr := gif.DecodeAll(upfile)
			if giferr != nil {
				res.Code = StatusImgDecode
				res.Msg = StatusText(StatusImgDecode)
				w.Write(ResponseJSON(res))
				return
			}
			err = gif.EncodeAll(file, gifimg)
		}

		if err != nil {
			res.Code = StatusImgEncode
			res.Msg = StatusText(StatusImgEncode)
			w.Write(ResponseJSON(res))
			return
		}
	}

	imgstr := fileMd5 + "_" + timeStamp + "." + imgtype
	res.Success = true
	res.Code = StatusOK
	res.Msg = StatusText(StatusOK)
	res.Data.Imgid = fileMd5 + "_" + timeStamp
	res.Data.Mime = imgtype
	res.Data.Size = upFileInfo.Size
	res.Data.ImgStr = imgstr

	// 打印上传成功日志
	log.Printf("Create file %s success\n", filePath)

	w.Write(ResponseJSON(res))
}
