package app

func DeleteApp(publicKey, dbURL string) bool {
	dbApp := &App{
		dbURL: dbURL,
		PublicKey: publicKey,
	}

	return dbApp.Delete()
}
