package config

import "github.com/fabricioque/gomysql/database"

// GetDBConection Conf
func GetDBConection() (config *database.DBConection) {
	return &database.DBConection{User: "root", Address: "127.0.0.1:3306", Schema: "gomysql"}
}
