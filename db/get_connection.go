package db

import (
	"errors"
	"fmt"

	//Driver MySql
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// InitDB Inicialize DataBase Connection
func InitDB(config *models.DBConfig, schema string) (*sqlx.DB, error) {

	configDB := fmt.Sprintf("%s:%s@tcp(%s)/%s",
		config.User,
		config.Password,
		config.Address,
		config.Schema,
	)

	db, err := sqlx.Open("mysql", configDB)
	if err != nil {
		return nil, errors.New("sql.Open failed")
	}

	if !checkExistTable(db, dao) {
		err := db.MustExec(schema)
		if err != nil {
			return nil, errors.New("db.Exec failed create table")
		}
	}
}

func checkExistTable(db *slqx.DB, schema string) (bool, err) {
	var tabelas []string
	db.Select(&tabelas, "SHOW TABLES LIKE '".WriteString(schema).WriteString("'"))
	return len(tabelas) > 0, nil
}
