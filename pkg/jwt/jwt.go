package jwt

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

// TokenExpireDuration 过期时间 7 天
const TokenExpireDuration = time.Hour * 24 * 7

// 设置签名密钥
var mySercet = []byte("qygzs uapply_go version1")

type MyClaims struct {
	// 自定义
	DepartmentID   int64  `json:"department_id"`
	OrganizationID int64  `json:"organization_id"`
	DepartmentName string `json:"department_name"`
	// 标准字段
	jwt.StandardClaims
}

// GenToken 生成JWT
func GenToken(organizationID, departmentID int64, departmentName string) (string, error) {
	// 创建一个我们自己的声明
	c := MyClaims{
		departmentID,
		// 自定义字段
		organizationID,
		departmentName,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(), // 过期时间
			Issuer:    "uapply_go",                                // 签发人
		},
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	return token.SignedString(mySercet)
}

// ParseToken 解析JWT
func ParseToken(tokenString string) (*MyClaims, error) {
	// 解析token
	var mc = new(MyClaims)
	token, err := jwt.ParseWithClaims(tokenString, mc, func(token *jwt.Token) (i interface{}, err error) {
		return mySercet, nil
	})
	if err != nil {
		return nil, err
	}
	if token.Valid { // 校验token
		return mc, nil
	}
	return nil, errors.New("invalid token")
}
