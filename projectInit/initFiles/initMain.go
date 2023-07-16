package initfiles

import (
	"os"
	"strings"

	"github.com/spf13/viper"
)

var mainFile = `package main

import (
	"{{PACKAGE_NAME}}/app"
)

// @title API {{NAME}}
// @version 0.1
// @description WRITE YOUR DESCRIPTION HERE.
// @contact.name WRITE YOUR NAME HERE.
// host localhost:3000
// @BasePath /
func main() {
	err := app.SetupAndRunApp()
	if err != nil {
		panic(err)
	}
}
`

func writeMainFile() string {
	packageName := viper.GetString("package_name")

	mainFile = strings.ReplaceAll(mainFile, "{{PACKAGE_NAME}}", packageName)
	mainFile = strings.ReplaceAll(mainFile, "{{NAME}}", viper.GetString("name"))

	return mainFile
}

func initMain() {
	filePath := "./main.go"

	if _, err := os.Stat(filePath); err == nil {
		return
	}

	file, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	_, err = file.WriteString(writeMainFile())
	if err != nil {
		panic(err)
	}
}
