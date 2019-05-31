package goliauth

import (
	"bytes"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateRandomKey(t *testing.T) {
	numbytes := 32
	key := NewRandomKey()

	assert.True(t, len(key) == numbytes, "must be 32 bytes")
	assert.True(t, len(fmt.Sprintf("%x", key)) == numbytes*2, "must be a 64 characters string")
	assert.False(t, bytes.Equal(key, make([]byte, numbytes)), "must be different than zeroes")
}

func TestEncryptDecrypt(t *testing.T) {

	secretKey := NewRandomKey()

	// set SECRET env variable
	os.Setenv("SECRET", fmt.Sprintf("%x", secretKey))

	appSecretKey := NewRandomKey()

	encryptedAppSecret, err := Encrypt(appSecretKey)

	if err != nil {
		t.Fatal(err)
	}

	decryptedAppSecret, err := Decrypt(encryptedAppSecret)

	if err != nil {
		t.Fatal(err)
	}

	assert.True(t, bytes.Equal(appSecretKey, decryptedAppSecret), "must decrypt to the same value")
}
