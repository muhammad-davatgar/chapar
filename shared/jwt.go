package shared

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type TokenType string

var (
	RefreshToken TokenType = "refresh_token"
	RefreshTime            = jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 7))
	AccessToken  TokenType = "access_token"
	AccessTime             = jwt.NewNumericDate(time.Now().Add(time.Hour * 24))
	ErrExpired             = fmt.Errorf("token expired")
)

type JWTCustomClaims struct {
	ID       uint   `json:"id"`
	UserName string `json:"user_name"`
	Type     TokenType
	jwt.RegisteredClaims
}

func NewClaims(id uint, userName string, tokenType TokenType) JWTCustomClaims {
	if tokenType == RefreshToken {
		return JWTCustomClaims{id, userName, tokenType, jwt.RegisteredClaims{ExpiresAt: RefreshTime}}
	} else {
		return JWTCustomClaims{id, userName, tokenType, jwt.RegisteredClaims{ExpiresAt: AccessTime}}

	}
}

func (c JWTCustomClaims) Valid() error {
	if c.ExpiresAt.Time.Unix() > time.Now().Unix() {
		return ErrExpired
	}

	return nil
}
