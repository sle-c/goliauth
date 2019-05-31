package goliauth

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseJWT(t *testing.T) {

	secret := "foobar"
	jwtToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.VfiYwvArp2lNV6UgpwgqrqfbJp9QpMdv07M8ZI4u4Vw"

	claims, err := ParseJWT(jwtToken, []byte(secret))

	require.NoError(t, err)
	assert.Equal(t, "John Doe", claims["name"], "must have name")
}

func TestEncodeJWT(t *testing.T) {
	secret := "foobar"
	claims := NewClaims(map[string]interface{}{
		"foo": "bar",
	})

	token, err := EncodeJWT(secret, claims)

	require.NoError(t, err)

	parsedClaims, err := ParseJWT(token, []byte(secret))

	require.NoError(t, err)
	assert.Equal(t, "bar", parsedClaims["foo"], "must be the same as encoded claims")
}
