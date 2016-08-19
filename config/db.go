package config

import "github.com/fabricioque/gomysql/db"

// GetDBConection Conf
func GetDBConection() (config *db.DBConection) {
	return &db.DBConection{User: "root", Address: "127.0.0.1:3306", Schema: "gomysql"}
}
