package models

// MySQLConfig contens struct create connection database.
type MySQLConfig struct {
	User     string
	Password string
	Address  string
	Schema   string
}
