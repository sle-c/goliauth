package app

import (
	"fmt"

	"github.com/omnisyle/goliauth"
)

func CreateApp(name, dbURL string) *App {
	publicKey := goliauth.NewRandomKey()
	secretKey := goliauth.NewRandomKey()
	encryptedSecret, err := goliauth.Encrypt(secretKey)

	if err != nil {
		panic(err)
	}

	app := &App{
		dbURL:           dbURL,
		Name:            name,
		PublicKey:       fmt.Sprintf("%x", publicKey),
		SecretKey:       fmt.Sprintf("%x", secretKey),
		EncryptedSecret: fmt.Sprintf("%x", encryptedSecret),
	}

	return app.Create()
}
