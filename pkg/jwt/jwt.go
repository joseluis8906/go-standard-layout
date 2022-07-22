package jwt

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	stdjwt "github.com/golang-jwt/jwt/v4"
	"github.com/joseluis8906/go-standard-layout/pkg/errors"
)

const (
	// ErrInvalidToken ...
	ErrInvalidToken = errors.Error("error invalid token")
)

// Encode ...
func Encode(ctx context.Context, secrect string, data interface{}, expireAt time.Duration) (string, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	claim := stdjwt.MapClaims{
		"exp": expireAt,
	}

	err = json.Unmarshal(jsonData, &claim)
	if err != nil {
		return "", err
	}

	token := stdjwt.NewWithClaims(stdjwt.SigningMethodHS256, claim)
	tokenString, err := token.SignedString([]byte(secrect))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// Decode ...
func Decode(ctx context.Context, secrect string, rawToken string, data interface{}) error {
	token, err := stdjwt.Parse(rawToken, func(token *stdjwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*stdjwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(secrect), nil
	})

	if err != nil {
		return err
	}

	claims, ok := token.Claims.(stdjwt.MapClaims)

	if ok && !token.Valid {
		return ErrInvalidToken
	}

	jsonData, err := json.Marshal(claims)
	if err != nil {
		return err
	}

	err = json.Unmarshal(jsonData, data)
	if err != nil {
		return err
	}

	return nil
}
