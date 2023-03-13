package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

// init is invoked before main()
func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

var (
	db *gorm.DB
)

func Connect() {
	// dsn := "host=localhost user=root password=secret dbname=simple_bank port=5432 sslmode=disable"
	postgresUrl, exists := os.LookupEnv("DATABASE_URL")
	if exists {
		fmt.Println(postgresUrl)
	}
	d, err := gorm.Open(postgres.Open(postgresUrl) , &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db = d
}

func GetBD() *gorm.DB {
	return db
}
