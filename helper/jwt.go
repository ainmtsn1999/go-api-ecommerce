package helper

import (
	"strconv"
	"time"

	"github.com/ainmtsn1999/go-api-ecommerce/config"
	"github.com/golang-jwt/jwt/v4"
)

type JwtCustomClaims struct {
	AuthId int    `json:"auth_id"`
	Email  string `json:"email"`
	jwt.StandardClaims
}

var (
	ttl    = config.GetEnvVariable("JWT_TTL_IN_MINUTES")
	secret = config.GetEnvVariable("JWT_SECRET")
)

func GenerateToken(authId int, authEmail string) (string, error) {
	jwtTtl, err := strconv.Atoi(ttl)
	if err != nil {
		panic(err)
	}

	claims := &JwtCustomClaims{
		AuthId: authId,
		Email:  authEmail,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(jwtTtl) * time.Minute).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
