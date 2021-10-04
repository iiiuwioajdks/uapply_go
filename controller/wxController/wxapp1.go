package wxController

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"log"
	"net/http"
	"uapply_go/logic/wxLogic"
	"uapply_go/response"
)

/*
 这个是报名者的小程序
*/

// Login1 微信小程序1登录
// @Summary 微信小程序1登录
// @Tags wxapp1
// @Accept application/json（接受数据类型）
// @Produce application/json （返回数据类型）
// @Security ApiKeyAuth
// @Success 200 {object} _Wx1Token
// @Param code query string true "wx.login 给的 code"
// @Failure 500
// @Router /wx1/login [get]
func Login1(c *gin.Context) {
	code := c.Query("code")
	token, err := wxLogic.Wxapp1Login(code)
	if err != nil {
		zap.L().Error("wxapp1 login error", zap.Error(err))
		log.Println(err)
		response.Fail(c, http.StatusInternalServerError, response.CodeSystemBusy)
		return
	}
	response.Success(c, token)
}
