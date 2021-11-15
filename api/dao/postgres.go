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
	return r.instance.Model(entity).Insert()
}

func (r *PostgresConnection) Get(id string, entity interface{}) (bool, error) {
	query := r.instance.Model(entity).Where("id = ?0", id)
	exists, _ := query.Exists()
	return exists, query.Select()
}

func (r *PostgresConnection) GetAll(entity interface{}) error {
	return r.instance.Model(entity).Select()
}

func (r *PostgresConnection) Update(entity interface{}) (interface{}, error) {
	return r.instance.Model(entity).WherePK().Update()
}

func (r *PostgresConnection) Delete(entity interface{}, id string) (interface{}, error) {
	query := r.instance.Model(entity).Where("id = ?0", id)
	query.Select()
	return query.Delete()
}
