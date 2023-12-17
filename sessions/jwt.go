package sessions

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const (
	secretKey          = "asdfasdafasdasdfasdf"
	expirationDuration = 1 * time.Minute
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

func Create(userID int64) (string, error) {
	currentTime := time.Now()
	expiresAt := currentTime.Add(expirationDuration)

	claim := Claim{
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(currentTime),
			ExpiresAt: jwt.NewNumericDate(expiresAt),
		},
		UserID: userID,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenStr, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	Register(tokenStr, expiresAt.Unix(), userID)

	return tokenStr, nil
}
