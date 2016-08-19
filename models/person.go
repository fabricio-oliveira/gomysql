package models

//Person Class models for sample
type Person struct {
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	Sex       bool   `db:"sex"`
}

// GetSchema schema create table databse
func (p *Person) GetSchema() string {
	return `CREATE TABLE person (
										first_name text,
										last_name text,
										sex bool
							);`
}

// GetNameDB name table databse
func (p *Person) GetNameDB() string { return "person" }
