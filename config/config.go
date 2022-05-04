package checkers

import (
	"log"
	"os"

	dotenv "github.com/joho/godotenv"
)

func init() {
	err := dotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
}

func get(key string) string {
	if os.Getenv(key) == "" {
		log.Fatal("Environment variable " + key + " is not set")
	}

	return os.Getenv(key)
}
