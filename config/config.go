package config

// db config
import (
	// "fmt"
	_ "github.com/joho/godotenv/autoload"
	"os"
)

var AppSecret = os.Getenv("APP_SECRET")

var DbName = os.Getenv("POSTGRES_DB")
var DbUser = os.Getenv("POSTGRES_USER")

// var DATABASE_URL = fmt.Sprintf("postgresql://%s@localhost:5432/%s?connect_timeout=5&sslmode=disable",
//   DbUser,
//   DbName)

var DATABASE_URL = "postgresql://health24@localhost:5432/health24_production?connect_timeout=5&sslmode=disable"
