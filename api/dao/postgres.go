package dao

import (
	"fmt"
	"github.com/go-pg/pg/v10"
	"os"
)

type PostgresConnection struct {
	instance *pg.DB
}

func (r *PostgresConnection) Connect() {
	if r.instance == nil {
		address := fmt.Sprintf("%s:%s", os.Getenv("HOST"), os.Getenv("PORT"))
		options := &pg.Options{
			User:     os.Getenv("USER"),
			Password: os.Getenv("PASSWORD"),
			Addr:     address,
			Database: os.Getenv("DB_NAME"),
			PoolSize: 50,
		}
		con := pg.Connect(options)
		if con == nil {
			fmt.Errorf("cannot connect to postgres")
		}
		r.instance = con
	}
}

func (r *PostgresConnection) Insert(entity interface{}) (interface{}, error) {
	// Executing SQL query for insertion
	return r.instance.Model(entity).Insert()
}

func (r *PostgresConnection) Get(id string, entity interface{}) error {
	// Executing query for single row
	return r.instance.Model(entity).Where(fmt.Sprintf("id = %s", id)).Select()
}
