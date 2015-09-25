package main

import (
	"fmt"

	"github.com/fabricioque/gomysql/models"
)

func main() {

	people := []models.Person{}
	New(Person).GetAll(&people)
	fmt.Println(people)
}
