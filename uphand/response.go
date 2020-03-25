package uphand

import "encoding/json"

// ResponseInterface 返回版本信息
type ResponseInterface interface {
	SetVersion(string)
}

// ResponseModel 响应请求的公共模型
type ResponseModel struct {
	Success bool   `json:"success"` // 是否成功
	Code    int    `json:"code"`    // 响应码
	Msg     string `json:"msg"`     // 响应信息
	Version string `json:"version"` // 版本号
	Data    string `json:"data"`    // 数据
}

// SetVersion 获取版本信息
func (r *ResponseModel) SetVersion(str string) {
	r.Version = str
}

// UpdateDate 上传响应数据
type UpdateDate struct {
	Size   int64  `json:"size"`   // 大小
	Mime   string `json:"mime"`   // 图片类型
	Imgid  string `json:"imgid"`  // 图片id
	ImgStr string `json:"imgstr"` // 带后缀的图片
}

// UpdateResponse 上传响应模型
type UpdateResponse struct {
	ResponseModel
	Data UpdateDate `json:"data"`
}

// ResponseJSON 响应 json 打包
func ResponseJSON(res ResponseInterface) []byte {

	res.SetVersion("v1.0.0")

	data, err := json.Marshal(res)
	if err != err {

		// 打包失败
		data, _ = json.Marshal(ResponseModel{false,
			StatusJSON,
			StatusText(StatusJSON),
			"", ""})
	}

	return data
}
