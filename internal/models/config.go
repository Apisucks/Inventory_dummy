package models

import (
	"encoding/json"
	"os"
)

var ConfigMap map[string]interface{}

func LoadConfig() {
	data, err := os.ReadFile("config.json")
	if err != nil {
		panic(err)
	}

	if err := json.Unmarshal(data, &ConfigMap); err != nil {
		panic(err)
	}
}
