package auth

import (
	"github.com/golang-jwt/jwt/v5"
	"fmt"
)

// VerifyToken parses and verifies the JWT, returning the token if it’s valid.
func ValidateToken(tokenString string) (*jwt.Token, error) {
    token, err := jwt.ParseWithClaims(tokenString, jwt.MapClaims{}, HS256KeyFunc)
    if err != nil {
        // Propagate the real error
        return nil, err
    }

    if !token.Valid {
        return nil, fmt.Errorf("invalid token")
    }

    // On success, return the parsed token
    return token, nil
}

func HS256KeyFunc (token *jwt.Token) (interface{}, error) {
	// Optional safety: ensure the signing method is what you expect
	if token.Method != jwt.SigningMethodHS256 {
		return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
	}
	return SecretKey, nil
}
