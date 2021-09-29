package jwt

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

// TokenExpireDuration Expiration time 7 days
const TokenExpireDuration = time.Hour * 24 * 7

// Set the signature key
var mySercet = []byte("qygzs uapply_go version1")

type MyClaims struct {
	// Customization
	DepartmentID   int64  `json:"department_id"`
	OrganizationID int64  `json:"organization_id"`
	DepartmentName string `json:"department_name"`
	// Standard field
	jwt.StandardClaims
}

// GenToken Build JWT
func GenToken(organizationID, departmentID int64, departmentName string) (string, error) {
	// Create our own statement
	c := MyClaims{
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

// ParseToken parse JWT
func ParseToken(tokenString string) (*MyClaims, error) {
	// parse token
	var mc = new(MyClaims)
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
