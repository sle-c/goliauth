package app

import (
	"context"
	"database/sql"
	"time"

	_ "github.com/lib/pq"
)

type App struct {
	ID              int    `db:"id"`
	Name            string `db:"name"`
	PublicKey       string `db:"public_key"`
	EncryptedSecret string `db:"encrypted_secret_key"`
	SecretKey       string
	CreatedAt       time.Time `db:"created_at"`
	UpdatedAt       time.Time `db:"updated_at"`

	dbURL string
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
		app.EncryptedSecret,
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

func (app *App) Get() *App {
	db := dialDB(app.dbURL)
	defer db.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*60)
	defer cancel()

	queryString := `SELECT id, name, encrypted_secret_key FROM apps where public_key = $1`
	row := db.QueryRowContext(
		ctx,
		queryString,
		app.PublicKey,
	)

	err := row.Scan(
		&app.ID,
		&app.Name,
		&app.EncryptedSecret,
	)

	if err != nil {
		panic(err)
	}

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
