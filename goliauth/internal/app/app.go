package app

import (
	"context"
	"database/sql"
	"time"

	_ "github.com/lib/pq"
)

type App struct {
	ID        int
	Name      string
	PublicKey string
	SecretKey string
	CreatedAt time.Time
	UpdatedAt time.Time

	encryptedSecret string
	dbURL           string
}

func (app *App) Create() *App {
	db := dialDB(app.dbURL)
	defer db.Close()

	now := time.Now().UTC()
	app.CreatedAt = now
	app.UpdatedAt = now

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*60)
	defer cancel()

	queryString := `
		INSERT INTO apps(
			name,
			public_key,
			encrypted_secret_key,
			created_at,
			updated_at
		) VALUES ($1, $2, $3, $4::timestamp, $5::timestamp)
		RETURNING id;
	`
	row := db.QueryRowContext(
		ctx,
		queryString,
		app.Name,
		app.PublicKey,
		app.encryptedSecret,
		app.CreatedAt,
		app.UpdatedAt,
	)
	var id int
	err := row.Scan(&id)

	if err != nil {
		panic(err)
	}

	app.ID = id

	return app
}

func dialDB(url string) *sql.DB {
	db, err := sql.Open("postgres", url)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db
}
