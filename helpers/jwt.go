package helpers

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const (
	secretKey = "asdfasdafasdasdfasdf"
)

type Claim struct {
	jwt.RegisteredClaims

	UserID int64
}

func GetUserIDFromJWT(token string) (int64, error) {
	var claim Claim
	if _, err := jwt.ParseWithClaims(token, &claim, func(t *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	}); err != nil {
		return 0, err
	}
	return claim.UserID, nil
}

func GenerateTokenFromUserID(userID int64) (string, error) {
	claim := Claim{
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
		},
		UserID: userID,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	return token.SignedString([]byte(secretKey))
}
