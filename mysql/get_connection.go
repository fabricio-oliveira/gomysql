package mysql

import (
	"errors"
	"fmt"
	"time"

	// driver mysl
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang/exemplo/gomysql/models"
	"github.com/jmoiron/sqlx"
)

// InitDB Inicialize DataBase Connection
func InitDB(config *models.MySQLConfig) (*sqlx.DB, error) {

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

	// var schema = `CREATE TABLE person (
	//            		first_name text,
	//            		last_name text,
	//      					gay bool
	//  						);`
	//
	// db.MustExec(schema)

	// if __, err = db.Exec(models.schema); err != nil {
	// 	return nil, errors.New("db.Exec failed create table")
	// }

	return db, nil
}

// ConvertTimeFromDB The name say everthing
func ConvertTimeFromDB(timeUnixNano int64) time.Time {
	t := int64(timeUnixNano) / int64(time.Second)
	return time.Unix(t, 0)
}
