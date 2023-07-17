package initfiles

import (
	"os"
	"strings"

	"github.com/spf13/viper"
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

var swaggerSetupFile = `package config

import (
	"{{IMPORT_API_LIB}}"
	"github.com/gofiber/swagger"
)

func AddSwaggerRoutes(app {{APP_TYPE}}) {
	// setup swagger
	app.{{GET_ROUTE}}("/swagger/*", swagger.HandlerDefault)
}
`

func writeSwaggerSetupFile() string {
	var replacement replacementApi

	if viper.GetString("api_library") == "Fiber" {
		replacement = FiberReplacement
	} else {
		replacement = GinReplacement
	}

	swaggerSetupFile = strings.ReplaceAll(swaggerSetupFile, "{{IMPORT_API_LIB}}", replacement.ImportApiLib)
	swaggerSetupFile = strings.ReplaceAll(swaggerSetupFile, "{{APP_TYPE}}", replacement.AppType)
	swaggerSetupFile = strings.ReplaceAll(swaggerSetupFile, "{{GET_ROUTE}}", replacement.GetRoute)

	return swaggerSetupFile
}

func initSetupEnv() {
	envFilePath := "./config/env.go"
	swaggerFilePath := "./config/swagger.go"

	if _, err := os.Stat(envFilePath); err == nil {
		return
	}
	if _, err := os.Stat(swaggerFilePath); err == nil {
		return
	}

	os.MkdirAll("./config", os.ModePerm)

	file, err := os.Create(envFilePath)
	if err != nil {
		panic(err)
	}

	_, err = file.WriteString(setupEnvFile)
	if err != nil {
		panic(err)
	}
	file.Close()

	file, err = os.Create(swaggerFilePath)
	if err != nil {
		panic(err)
	}

	_, err = file.WriteString(writeSwaggerSetupFile())
	if err != nil {
		panic(err)
	}
	file.Close()
}
