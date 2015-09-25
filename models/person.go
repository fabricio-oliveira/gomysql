package models

import (
	"errors"

	"github.com/fabricioque/gomysql/mysql"
	"github.com/jmoiron/sqlx"
)

//Person Class models for sample
type Person struct {
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	Gay       bool   `db:"gay"`
	db        *sqlx.DB
}

// GetSchema databse go
func getSchema() string {
	return `CREATE TABLE person (
										first_name text,
										last_name text,
										gay bool
							);`
}

// GetAll Person DataBase
func (p *Person) GetAll() (*Person, error) {
	if db == nil {
		db, err := mysql.InitDB(config, getSchema())
		if err == nil {
			return nil, errors.New("mysql.InitDB failed")
		}
	}

	if err = db.Select(&people, "select * from person;"); err != nil {
		return nil, errors.New("sql.Select failed")
	}

	return people, nil
}
