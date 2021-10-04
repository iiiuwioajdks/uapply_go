package wxController

import (
	"uapply_go/response"
)

type _Wx1Token struct {
	Code    response.ResCode `json:"code"` // 业务响应状态码
	Message string           `json:"msg"`  // 提示信息
	Data    string           `json:"data"` // token
}
