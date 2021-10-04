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
