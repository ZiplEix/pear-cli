package initfiles

import (
	"os"
)

var setupEnvFile = `package config

import (
	"os"

	"github.com/joho/godotenv"
)

func LoadENV() error {
	goEnv := os.Getenv("GO_ENV")

	if goEnv == "" || goEnv == "development" {
		err := godotenv.Load()
		if err != nil {
			return err
		}
	}
	return nil
}
`

func initSetupEnv() {
	filePath := "./config/env.go"

	if _, err := os.Stat(filePath); err == nil {
		return
	}

	os.MkdirAll("./config", os.ModePerm)

	file, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	_, err = file.WriteString(setupEnvFile)
	if err != nil {
		panic(err)
	}
}
