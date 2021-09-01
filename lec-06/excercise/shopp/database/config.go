// package database

// import(
// 	"fmt"
// )

// type Config struct{
// 	ServerName string
// 	Username string
// 	Password string
// 	DB string
// }

// func GetConnectionString(conf *Config)  string{
// 	connectionString := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&collation=utf8mb4_unicode_ci&parseTime=true&multiStatements=true", config.Username, config.Password, config.ServerName, config.DB)

// 	return connectionString
// }

package database

import (
	"fmt"
)

type Config struct{

	Server string
	Username string
	Password string
	DB string
}


func GetConnectionString (config Config)  string{
	connectionString := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&collation=utf8mb4_unicode_ci&parseTime=true&multiStatements=true", config.Username, config.Password, config.Server, config.DB)

	return connectionString
}

