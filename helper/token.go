package helper

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jago-bank-api/entity"
	"time"
)

var mySingingKey = []byte("secret")

type JWTClaims struct {
	ID int `json:"id"`
	jwt.RegisteredClaims
}

func GenerateToken(user *entity.User) (string, error) {
	claims := JWTClaims{
		int(user.ID),
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(60 * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySingingKey)

	return ss, err
}

func ValidateToken(tokenString string) (*int, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return mySingingKey, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return nil, errors.New("Invalid token signature")
		}

		return nil, errors.New("Your token is expired")
	}

	claims, ok := token.Claims.(*JWTClaims)
	if !ok || !token.Valid {
		return nil, errors.New("Your token is expired")
	}

	return &claims.ID, nil
}
