package common

import (
	"time"

	"github.com/Yingzhixin/gin_demo/models"
	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("a_secret_crect")

//Claims 用户验证
type Claims struct {
	UserID uint
	jwt.StandardClaims
}

//ReleaseToken 建立token
func ReleaseToken(user models.User) (string, error) {
	//设置token过期时间
	expirationTime := time.Now().Add(300 * time.Second)

	claims := &Claims{
		UserID: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "yingzhixin",
			Subject:   "user token",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		return "", err
	}
	return tokenString, nil
}

//ParseToken 从token中解析出claim
func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	return token, claims, err
}
