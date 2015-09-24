package models

// Person Class models for sample
type Person struct {
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	Gay       bool   `db:"gay"`
}

// GetSchema databse go
func GetSchema() string {
	return `CREATE TABLE person (
										first_name text,
										last_name text,
										gay bool
							);`
}
