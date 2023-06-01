package utils

import (
	"fmt"

	"github.com/golang-jwt/jwt"
)

func ParseToken(token string, signedJWTKey string) (string, string, error) {
	tok, err := jwt.Parse(token, func(jwtToken *jwt.Token) (interface{}, error) {
		if _, ok := jwtToken.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected method: %s", jwtToken.Header["alg"])
		}

		return []byte(signedJWTKey), nil
	})

	if err != nil {
		return "", "", fmt.Errorf("invalidate token: %w", err)
	}

	claims, ok := tok.Claims.(jwt.MapClaims)
	if !ok || !tok.Valid {
		return "", "", fmt.Errorf("invalid token claim")
	}

	return claims["role"].(string), claims["uid"].(string), nil
}
