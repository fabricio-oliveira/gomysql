package main

import (
	"fmt"

	"github.com/fabricioque/gomysql/config"
	"github.com/fabricioque/gomysql/models"
)

func main() {

	people := []models.Person{}
	connection := config.GetDBConection()

	connection.GetAll(&people)
	fmt.Println(people)
}
