package dao

import (
	"errors"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type PostgresConnection struct {
	instance   *gorm.DB
	connString string
}

func NewPostgresDB(host, port, user, password, dbname string) *PostgresConnection {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, port)
	return &PostgresConnection{
		connString: dsn,
	}
}

func (r *PostgresConnection) Connect() {
	if r.instance == nil {
		db, err := gorm.Open(postgres.Open(r.connString), &gorm.Config{})
		if err != nil {
			fmt.Errorf("cannot connect to postgres")
		}
		r.instance = db
	}
}

func (r *PostgresConnection) Insert(entity interface{}) (interface{}, error) {
	result := r.instance.Create(entity)
	return result, result.Error
}

func (r *PostgresConnection) Get(id string, entity interface{}, preload ...string) (bool, error) {
	var result *gorm.DB
	if len(preload) != 0 {
		result = r.instance.Preload(preload[0]).First(entity, id)
	} else {
		result = r.instance.First(entity, id)
	}
	return !errors.Is(result.Error, gorm.ErrRecordNotFound), result.Error
}

func (r *PostgresConnection) GetAll(entity interface{}) error {
	return r.instance.Find(entity).Error
}

func (r *PostgresConnection) Update(entity interface{}) (interface{}, error) {
	result := r.instance.Save(entity)
	return result, result.Error
}

func (r *PostgresConnection) Delete(entity interface{}, id string) (interface{}, error) {
	result := r.instance.Clauses(clause.Returning{}).Delete(entity, id)
	return result, result.Error
}
