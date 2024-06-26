package initializers

import "github.com/joho/godotenv"

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		panic("Failed to load .env file!")
	}
}
