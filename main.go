package main

import (
	"fmt"
	"log"

	"github.com/golang/exemplo/gomysql/models"
	"github.com/golang/exemplo/gomysql/mysql"
)

func main() {
	config := &models.MySQLConfig{"root", "", "127.0.0.1:3306", "gomysql"}
	db, err := mysql.InitDB(config)
	if err != nil {
		log.Fatal(err)
	}

	people := []models.Person{}
	query := "select * from person;"
	if err = db.Select(&people, query); err != nil {
		log.Fatal(err)
	}
	fmt.Println(people)
}
