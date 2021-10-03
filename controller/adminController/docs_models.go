package adminController

import (
	"uapply_go/entity/DBModels"
	"uapply_go/response"
)

type _Organizations struct {
	Code    response.ResCode          `json:"code"`    // 业务响应状态码
	Message string                    `json:"message"` // 提示信息
	Data    []*DBModels.Organizations `json:"data"`    // 返回token
}
