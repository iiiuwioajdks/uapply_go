package testController

import (
	"github.com/gin-gonic/gin"
	"uapply_go/entity/ResponseModels"
	"uapply_go/logic/testLogic"
	"uapply_go/response"
)

// Pong pingpong测试（Pong controller 的作用）
// @Summary pingpong测试（总结）
// @Description 测试的描述信息
// @Tags test（标签）
// @Accept application/json（接受数据类型）
// @Produce application/json （返回数据类型）
// @Param Authorization header string false "Bearer 用户令牌" （token类型）
// @Security ApiKeyAuth
// @Success 200 {object} _Pong
// @Router /ping [get]
func Pong(c *gin.Context) {
	var p ResponseModels.Pong
	err := testLogic.Pong(&p)
	if err != nil {

	}
	response.Success(c, p)
}
