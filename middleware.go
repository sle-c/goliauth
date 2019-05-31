package goliauth

import (
	"fmt"
	"net/http"
	"strings"
)

// AuthenticateJWTMiddleware wraps AuthenticateJWTToken to provide middleware
// this is just an example to show how it can be used as a middleware
func AuthenticateJWTMiddleware(next http.Handler, secretKey string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := AuthenticateJWTToken(secretKey, r)

		if err == nil {
			next.ServeHTTP(w, r)
		} else {
			http.Error(w, "Forbidden", http.StatusForbidden)
		}
	})
}

// AuthenticateJWTToken is the main function to
// verify the JWT token from a request and it returns the claims
func AuthenticateJWTToken(secretKey string, req *http.Request) (map[string]interface{}, error) {
	jwtToken, err := ExtractJWTToken(req)

	if err != nil {
		return nil, fmt.Errorf("Failed get JWT token")
	}

	claims, err := ParseJWT(jwtToken, []byte(secretKey))

	if err != nil {
		return nil, fmt.Errorf("Failed to parse token")
	}

	return claims, nil
}

// ExtractJWTToken extracts bearer token from Authorization header
func ExtractJWTToken(req *http.Request) (string, error) {
	tokenString := req.Header.Get("Authorization")

	if tokenString == "" {
		return "", fmt.Errorf("Could not find token")
	}

	tokenString, err := stripTokenPrefix(tokenString)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// Strips 'Token' or 'Bearer' prefix from token string
func stripTokenPrefix(tok string) (string, error) {
	// split token to 2 parts
	tokenParts := strings.Split(tok, " ")

	if len(tokenParts) < 2 {
		return tokenParts[0], nil
	}

	return tokenParts[1], nil
}
