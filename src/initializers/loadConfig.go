package initializers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"main/src/config"
	"os"
)

func LoadConfig() {
	jsonFile, err := os.Open("config.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal(byteValue, &config.Configuration)
}
