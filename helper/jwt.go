package helper

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/ainmtsn1999/go-api-ecommerce/config"
	"github.com/golang-jwt/jwt/v4"
)

type Token struct {
	AuthId int    `json:"auth_id"`
	Email  string `json:"email"`
}

var (
	ttl    = config.GetEnvVariable("JWT_TTL_IN_MINUTES")
	secret = config.GetEnvVariable("JWT_SECRET")
)

func GenerateToken(payload *Token) (string, error) {
	jwtTtl, err := strconv.Atoi(ttl)
	if err != nil {
		panic(err)
	}

	claims := jwt.MapClaims{
		"payload": payload,
		"issued":  time.Now().Add(time.Duration(jwtTtl) * time.Minute),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyToken(tokenString string) (*Token, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}

		return []byte(secret), nil
	})

	if err != nil {
		return nil, err
	}

	// Check valid token
	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid token")

	}

	// Validate issued token
	issuedString := fmt.Sprintf("%v", claims["issued"])
	issued, err := time.Parse(time.RFC3339, issuedString)
	if err != nil {
		return nil, err
	}

	if time.Now().After(issued) {
		return nil, errors.New("token expired")
	}

	byteClaims, err := json.Marshal(claims["payload"])
	if err != nil {
		return nil, err
	}

	var myToken Token
	err = json.Unmarshal(byteClaims, &myToken)
	if err != nil {
		return nil, err
	}

	return &myToken, nil
}
