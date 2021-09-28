package departmentController

import (
	"uapply_go/response"
)

type _Token struct {
	Code    response.ResCode `json:"code"`    // 业务响应状态码
	Message string           `json:"message"` // 提示信息
	Data    string           `json:"data"`    // 返回token
}

type _TokenFail struct {
	Code    response.ResCode `json:"code"`    // 业务响应状态码
	Message string           `json:"message"` // 提示信息
}
