package config

// db config
import (
	"fmt"
	_ "github.com/joho/godotenv/autoload"
	"os"
)

var AppSecret = os.Getenv("APP_SECRET")

// determine project ENV
var GO_ENV = os.Getenv("GO_ENV")

func EnvProduction() bool {
	return GO_ENV == "PRODUCTION"
}

func EnvDevelopment() bool {
	return GO_ENV == "DEVELOPMENT" || GO_ENV == ""
}

// DB connection
var DbName = os.Getenv("POSTGRES_DB")
var DbUser = os.Getenv("POSTGRES_USER")

var DATABASE_URL = fmt.Sprintf("postgresql://%s@localhost:5432/%s?connect_timeout=5&sslmode=disable",
	DbUser,
	DbName)
