package constants

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func ReadKey(which string) (token string) {
	Keys := Keys{}

	jsonFile, err := os.Open("config/config.json")
	if err != nil {
		fmt.Println("Deu ruim")
	}

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println("Deu ruim")
	}

	err = json.Unmarshal(byteValue, &Keys)
	if err != nil {
		fmt.Println("Deu ruim")
	}

	if which == "BotKey" {
		token = Keys.BotKey
	} else {
		token = Keys.NewsKey
	}

	return token
}

type Keys struct {
	BotKey  string `json:"botKey"`
	NewsKey string `json:"newsKey"`
}
