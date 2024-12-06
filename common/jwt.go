package common

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// JwtPayLoad jwt中payload数据
type JwtPayLoad struct {
	UserID   uint   `json:"userId"`   // 用户id
	Username string `json:"username"` // 用户名
	Address  string `json:"address"`  // 用户绑定地址
}

// Custom claims (可以根据需要扩展)
type CustomClaims struct {
	JwtPayLoad
	jwt.RegisteredClaims
}

// GenerateToken 生成 JWT
func GenerateToken(user JwtPayLoad, accessSecret string, expires int64) (string, error) {
	claims := CustomClaims{
		JwtPayLoad: user,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(expires))),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "doge-sniper", // 发放者
		},
	}

	// 使用 HMAC SHA256 签名
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(accessSecret))
}

// ValidateToken 验证 JWT 是否有效
func ValidateToken(tokenString string, accessSecret string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		// 确认使用的签名方法
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return accessSecret, nil
	})

	// 验证失败
	if err != nil {
		return nil, err
	}

	// 解析 Token 中的 Claims
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}
