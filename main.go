package main

import (
	"fmt"

	"github.com/fabricioque/gomysql/database"
	"github.com/fabricioque/gomysql/models"
)

func main() {

	people := []models.Person{}
	connection := database.DBConection{"root", "", "127.0.0.1:3306", "gomysql"}

	connection.GetAll(&people)
	fmt.Println(people)
}
