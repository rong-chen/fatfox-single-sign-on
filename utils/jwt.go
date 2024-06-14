package utils

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"time"
)

var (
	AccessToken  = "access_token"
	RefreshToken = "refresh_token"
)

type Params struct {
	Id       uuid.UUID `gorm:"primaryKey"`
	Username string    `json:"username" form:"username" gorm:"username" `
	Phone    string    `json:"phone"  form:"phone" gorm:"phone"`
	Email    string    `json:"email;not null;unique"  form:"email" gorm:"email"`
	Nickname string    `json:"nickname"  form:"nickname" gorm:"nickname"`
	Avatar   string    `json:"avatar"  form:"avatar" gorm:"avatar"`
}

type Token struct {
	Params               Params `json:"user"`
	Type                 string `json:"type"`
	jwt.RegisteredClaims        // v5版本新加的方法
}

var secretKey = []byte("user-registration-center")

func GenerateJWT(params Params, types string, ex time.Time) (string, error) {
	token := Token{
		params,
		types,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(ex), // 30天过期
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, token)

	resp, err := t.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return resp, nil
}

func ParseJWT(token string) (*Token, error) {
	t, err := jwt.ParseWithClaims(token, &Token{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	fmt.Println(t)
	if claims, ok := t.Claims.(*Token); ok && t.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}
