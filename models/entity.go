package models

//Entity is a Interface for DAO Object
type Entity interface {
	GetSchema() string
	GetNameDB() string
}
