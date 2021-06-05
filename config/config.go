package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Configuration struct {
	BaseURL string
	Author  string
}

var (
	Config Configuration
)

func init() {
	file, err := os.Open("static/config.json")
	if err != nil {
		fmt.Println("Couldn't find static/config.json")
		os.Exit(1)
	}

	defer file.Close()

	parser := json.NewDecoder(file)
	err = parser.Decode(&Config)
	if err != nil {
		fmt.Println("Malformed JSON in configuration file")
		os.Exit(1)
	}
}
