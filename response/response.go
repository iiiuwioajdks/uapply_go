package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Success Encapsulated the successful return code
func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code": CodeSuccess,
		"msg":  CodeSuccess.Msg(),
		"data": data,
	})
}

// Fail Normal failure returns
func Fail(c *gin.Context, status int, code ResCode) {
	c.JSON(status, gin.H{
		"code": code,
		"msg":  code.Msg(),
	})
}

// FailWithMsg Return with an error message,
// note that the code here cannot be a ResCode type, that is, it cannot be defined in code
func FailWithMsg(c *gin.Context, status int, code int64, msg interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg,
	})
}
