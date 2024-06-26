package initializers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"main/src/config"
	"main/src/utilities"

	"github.com/joho/godotenv"
)

func Initialize(args []string) {
	loadEnv()
	loadConfig()
	utilities.DetectMigration()
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		panic("Failed to load .env file!")
	}
}

func loadConfig() {
	jsonFile, err := os.Open("config.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal(byteValue, &config.Configuration)
}
