package jwt

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
	"uapply_go/entity/ResponseModels"
)

// TokenExpireDuration Expiration time 7 days
const TokenExpireDuration = time.Hour * 24 * 7

// Set the signature key
var mySercet = []byte("qygzs uapply_go version1")

type WebClaims struct {
	// Customization
	DepartmentID   int64  `json:"department_id"`
	OrganizationID int64  `json:"organization_id"`
	DepartmentName string `json:"department_name"`
	// Standard field
	jwt.StandardClaims
}

type WxAppClaims struct {
	SessionKey string `json:"session_key"`
	OpenID     string `json:"openid"`
	jwt.StandardClaims
}

// GenToken Build Web JWT
func GenToken(organizationID, departmentID int64, departmentName string) (string, error) {
	// Create our own statement
	c := WebClaims{
		departmentID,
		// Customization
		organizationID,
		departmentName,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(), // 过期时间
			Issuer:    "uapply_go",                                // 签发人
		},
	}
	// Create a signature object using the SigningMethodHS256 signature method
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// Use the specified seek signature and get the fully encoded token
	return token.SignedString(mySercet)
}

func GenToken2(ws1 *ResponseModels.WxSession1) (string, error) {
	ws1.ExpireIn = time.Now().Add(TokenExpireDuration).Unix()
	c := WxAppClaims{
		SessionKey: ws1.SessionKey,
		OpenID:     ws1.OpenID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: ws1.ExpireIn,
			Issuer:    "uaaply_go",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return token.SignedString(mySercet)
}

// ParseToken parse web JWT
func ParseToken(tokenString string) (*WebClaims, error) {
	// parse token
	var mc = new(WebClaims)
	token, err := jwt.ParseWithClaims(tokenString, mc, func(token *jwt.Token) (i interface{}, err error) {
		return mySercet, nil
	})
	if err != nil {
		return nil, err
	}
	if token.Valid {
		return mc, nil
	}
	return nil, errors.New("invalid token")
}

func ParseToken2(tokenString string) (*WxAppClaims, error) {
	// parse token
	var mc = new(WxAppClaims)
	token, err := jwt.ParseWithClaims(tokenString, mc, func(token *jwt.Token) (i interface{}, err error) {
		return mySercet, nil
	})
	if err != nil {
		return nil, err
	}
	if token.Valid {
		return mc, nil
	}
	return nil, errors.New("invalid token")
}
