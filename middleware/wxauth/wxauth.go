package wxauth

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"uapply_go/pkg/jwt"
	"uapply_go/response"
)

/*
微信小程序登录 token认证中间件
*/

const OpenIDKey = "openid"
const SessionIDKey = "session_id"

// Wx1JWTAuthMiddleware JWT-based certified middleware
func Wx1JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		// 客户端携带Token有三种方式 1.放在请求头 2.放在请求体 3.放在URI
		// 这里假设Token放在Header的Authorization中，并使用Bearer开头
		// 这里的具体实现方式要依据你的实际业务情况决定
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			response.Fail(c, http.StatusUnauthorized, response.CodeTokenInvalid)
			c.Abort()
			return
		}
		// Split by space
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			response.Fail(c, http.StatusUnauthorized, response.CodeTokenInvalid)
			c.Abort()
			return
		}
		// parts[1] is the acquired tokenString,
		//and we use the previously defined function to parse JWT to parse it
		ms, err := jwt.ParseToken2(parts[1])
		if err != nil {
			response.Fail(c, http.StatusUnauthorized, response.CodeTokenInvalid)
			c.Abort()
			return
		}
		// Save the currently requested message information to the requested context c
		c.Set(OpenIDKey, ms.OpenID)
		c.Set(SessionIDKey, ms.SessionKey)
		// 需要用到 openid的时候直接拿就行了
		c.Next()
	}
}
