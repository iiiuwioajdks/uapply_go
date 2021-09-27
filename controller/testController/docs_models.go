package testController

/*
用来存放 swagger 的models
*/
import (
	"uapply_go/entity/ResponseModels"
	"uapply_go/response"
)

type _Pong struct {
	Code    response.ResCode     `json:"code"`    // 业务响应状态码
	Message string               `json:"message"` // 提示信息
	Data    *ResponseModels.Pong `json:"data"`    // 数据
}
