package util

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"log"
	"time"
)

// MyClaims 自定义的payload
type MyClaims struct {
	jwt.RegisteredClaims        // 内置的payload字段
	UserID               string `json:"user_id"` // 自定义的payload字段
}

type JWT struct {
}

var key = []byte("qwouidfoqiwr23fioqw") // 用于签名的密钥

func (j *JWT) GenerateToken(msg string, ttl int64) string {
	now := time.Now()
	// 创建一个新的token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &MyClaims{
		UserID: msg,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(now.Add(time.Duration(24*ttl) * time.Hour)), // 过期时间，ttl天后过期
			IssuedAt:  jwt.NewNumericDate(now),                                        // 签发时间
			NotBefore: jwt.NewNumericDate(now),                                        // 生效时间
			Issuer:    "tiktok-lite",                                                  // 签发人
		},
	})
	tokenString, _ := token.SignedString(key) // 签名, 返回token字符串
	return tokenString
}

func (j *JWT) ParseToken(tokenString string) (string, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})

	if err != nil {
		log.Println(err)
		return "", errors.New("invalid token")
	}

	claims, ok := token.Claims.(*MyClaims)
	if !ok || !token.Valid {
		log.Println("invalid token")
		return "", errors.New("invalid token")
	}
	return claims.UserID, nil
}
