package auth

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const (
	expiresIn = time.Hour
)

var (
	secret = "secret"
)

func SetSecret(key string) {
	secret = key
}

// 😇😇😇
func MakeJWT(id string) (string, error) {
	secretKey := []byte(secret)

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = id
	claims["exp"] = time.Now().Add(expiresIn).Unix()

	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		fmt.Println("Error creating token:", err)
		return "", err
	}

	return signedToken, nil
}

var (
	ErrTokenExpired = errors.New("token expired")
	ErrTokenInvalid = errors.New("token invalid")
)

// 😇😇😇
func DecodeJWT(jwtToken string) (string, error) {
	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil {
		return "", ErrTokenInvalid
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		name := claims["name"].(string)
		expiration := int64(claims["exp"].(float64))
		if expiration > time.Now().Unix() {
			return name, nil
		} else {
			return "", ErrTokenExpired
		}
	} else {
		return "", ErrTokenInvalid
	}
}
