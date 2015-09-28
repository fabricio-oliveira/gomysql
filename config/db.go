package config

import "github.com/fabricioque/gomysql/database"

// GetDBConfig Conf
func GetDBConfig() (config *database.DBConfig) {
	return &models.DBConfig{"root", "", "127.0.0.1:3306", "gomysql"}
}
