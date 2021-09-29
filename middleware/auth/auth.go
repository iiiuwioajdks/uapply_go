package auth

import (
	"github.com/gin-gonic/gin"
	"strings"
	"uapply_go/pkg/jwt"
)

const OrganizationIdKey = "organization_id"
const DepartmentIdKey = "department_id"

// JWTAuthMiddleware JWT-based certified middleware
func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		// 客户端携带Token有三种方式 1.放在请求头 2.放在请求体 3.放在URI
		// 这里假设Token放在Header的Authorization中，并使用Bearer开头
		// 这里的具体实现方式要依据你的实际业务情况决定
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			c.Abort()
			return
		}
		// Split by space
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {

			c.Abort()
			return
		}
		// parts[1] is the acquired tokenString,
		//and we use the previously defined function to parse JWT to parse it
		mc, err := jwt.ParseToken(parts[1])
		if err != nil {
			c.Abort()
			return
		}
		// Save the currently requested message information to the requested context c
		c.Set(OrganizationIdKey, mc.OrganizationID)
		c.Set(DepartmentIdKey, mc.DepartmentID)
		// Subsequent handlers can use c.Get(OrganizationIdKey) gets the organization
		// and department information that is currently requested
		c.Next()
	}
}
