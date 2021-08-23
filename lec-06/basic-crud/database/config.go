package database

import "fmt"

type Config struct{

	ServerName string
	Username string
	Password string
	DB string
}


var s = 4;
var GetConnectionString = func (config Config)  string{
	connectionString := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&collation=utf8mb4_unicode_ci&parseTime=true&multiStatements=true", config.Username, config.Password, config.ServerName, config.DB)

	return connectionString
}