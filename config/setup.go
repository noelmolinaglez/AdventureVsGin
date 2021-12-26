package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {

	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("Not exist enviroment.")
	}

	server := os.Getenv("SERVER")
	port := os.Getenv("DATABASEPORT")
	database := os.Getenv("DATABASE")
	user := os.Getenv("DATABASEUSER")
	password := os.Getenv("PASSWORD")

	dsn := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s", user, password, server, port, database)
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Wrong connection.")
	}

	DB = db
}
