package main

import (
	"fmt"

	"github.com/fabricioque/gomysql/config"
)

func main() {

	connection := config.GetDBConection()
	connection.Init()

	people, _ := connection.GetAll()
	fmt.Println(people)
}
