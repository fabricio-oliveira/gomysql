package database

import (
	"errors"
	"fmt"

	"github.com/fabricioque/gomysql/models"
	"github.com/jmoiron/sqlx"
)

// DBConection contens struct create connection database.
type DBConection struct {
	User     string
	Password string
	Address  string
	Schema   string
	db       *sqlx.DB
}

// InitT Inicialize DataBase Connection
func (conection *DBConection) InitT() error {

	configDB := fmt.Sprintf("%s:%s@tcp(%s)/%s",
		conection.User,
		conection.Password,
		conection.Address,
		conection.Schema,
	)
	var err error
	conection.db, err = sqlx.Open("mysql", configDB)
	if err != nil {
		return errors.New("sql.Open failed")
	}

	return nil
}

func (conection *DBConection) checkExistTable(schema string) (bool, error) {
	var tabelas []string

	if conection.db == nil {
		return false, errors.New("DB.checkExistTable conection db not exist")
	}

	query := fmt.Sprintf("SHOW TABLES LIKE %s'", schema)
	err := conection.db.Select(&tabelas, query)
	if err != nil {
		return false, err
	}

	return len(tabelas) > 0, nil
}

// GetAll Person DataBase
func (conection *DBConection) GetAll(m *[]models.Entity) error {

	if conection.db == nil {
		return errors.New("DB.GetAll conection not open")
	}

	query := fmt.Sprintf("select * from %s", (*m)[0].GetNameDB())
	if err := conection.db.Select(m, query); err != nil {
		return errors.New("sql.Select failed")
	}
	return nil
}
