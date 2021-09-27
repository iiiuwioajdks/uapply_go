package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Success 封装成功的返回码
func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code": CodeSuccess,
		"msg":  CodeSuccess.Msg(),
		"data": data,
	})
}

// Fail 普通失败返回
func Fail(c *gin.Context, code ResCode) {
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  code.Msg(),
	})
}

// FailWithMsg 带着错误信息返回，注意这里的 code 不能是 ResCode类型，也就是不能在code里面定义
func FailWithMsg(c *gin.Context, code int64, msg interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg,
	})
}
