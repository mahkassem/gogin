package initializers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"main/src/config"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		panic("Failed to load .env file!")
	}
}

func LoadConfig() {
	jsonFile, err := os.Open("config.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal(byteValue, &config.Configuration)
}
