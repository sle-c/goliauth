package app

import (
	"encoding/hex"
	"fmt"

	"github.com/omnisyle/goliauth"
)

func GetApp(publicKey, dbURL string) *App {
	app := &App{
		PublicKey: publicKey,
		dbURL:     dbURL,
	}

	dbApp := app.Get()

	secretBytes, err := hex.DecodeString(dbApp.EncryptedSecret)

	if err != nil {
		panic(err)
	}

	decryptedKey, err := goliauth.Decrypt(secretBytes)
	if err != nil {
		panic(err)
	}

	dbApp.SecretKey = fmt.Sprintf(
		"%x",
		decryptedKey,
	)

	return dbApp
}
