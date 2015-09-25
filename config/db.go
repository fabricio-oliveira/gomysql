package config

import "github.com/fabricioque/gomysql/models"

// GetDBConfig Conf
func GetDBConfig() (config *models.DBConfig) {
	return &models.DBConfig{"root", "", "127.0.0.1:3306", "gomysql"}
}
