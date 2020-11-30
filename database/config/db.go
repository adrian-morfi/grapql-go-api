package config

import "fmt"

const (
	DBUser     = "postgres"
	DBPassword = "postgres"
	DBName     = "testdb"
	DBHost     = "127.0.0.1"
	DBPort     = "5555"
	DBType     = "postgres"
)

func GetDBType() string {
	return DBType
}

func GetPostgresConnectionString() string {
	dataBase := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		DBHost,
		DBPort,
		DBUser,
		DBName,
		DBPassword)
	return dataBase
}
