package dao

import "os"

func CreateEmailerDB() DBConnection {
	return NewDBConnection(NewPostgresDB(
		os.Getenv("HOST"),
		os.Getenv("PORT"),
		os.Getenv("USER"),
		os.Getenv("PASSWORD"),
		os.Getenv("DB_NAME"),
	))
}
