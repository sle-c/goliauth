package goliauth

import (
	"fmt"

	jwt "github.com/dgrijalva/jwt-go"
)

// Claims is an alias for MapClaims
type Claims = jwt.MapClaims

// StandardClaims wraps jwt standard claims type
type StandardClaims jwt.StandardClaims

// NewClaims create a Claims type
func NewClaims(data map[string]interface{}) Claims {
	newClaims := Claims(data)
	return newClaims
}

// // Valid implements the jwt.Claims interface
// func (cl Claims) Valid() error {
// 	return nil
// }

// ParseJWT parses a JWT and returns Claims object
// Claims can be access using index notation such as claims["foo"]
func ParseJWT(tokenString string, key []byte) (Claims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return key, nil
	})

	if token.Valid {
		if claims, ok := token.Claims.(Claims); ok {
			return claims, nil
		}

		return nil, err
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			return nil, err
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			// Token is either expired or not active yet
			return nil, err
		} else {
			return nil, err
		}
	} else {
		return nil, err
	}
}

// EncodeJWT serialize data into a jwt token using a secret
// This secret must match with the client's secret who's generating the token
func EncodeJWT(secretKey string, claims Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := []byte(secretKey)
	tokenString, err := token.SignedString(secret)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}
