package initfiles

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/viper"
)

var setupFile = `package app

import (
	"os"

	"{{PACKAGE_NAME}}/config"
	"{{PACKAGE_NAME}}/database"
	"{{PACKAGE_NAME}}/router"
	"{{IMPORT_API_LIB}}"
	"{{IMPORT_API_LIB_MIDDLEWARE_CORS}}"
	"{{IMPORT_API_LIB_MIDDLEWARE_LOGGER}}"
	"{{IMPORT_API_LIB_MIDDLEWARE_RECOVER}}"
)

func SetupAndRunApp() error {
	// Load ENV
	err := config.LoadENV()
	if err != nil {
		return err
	}

	// start database
	err = database.ConnectDB()
	if err != nil {
		return err
	}

	// defer database close
	defer database.CloseDB()

	// create app
	app := {{APP_CREATION}}

	// set log for api lib
	app.Use({{API_LOGGER}})
	app.Use({{API_CORS}})

	// recover from panic
	app.Use(recover.New())

	// setup routes
	router.SetupRoutes(app)

	// attach swagger
	config.AddSwaggerRoutes(app)

	// start app
	port := os.Getenv("PORT")
	app.{{API_LISTEN}}(":" + port)

	return nil
}
`

func writeSetupFile() string {
	packageName := viper.GetString("package_name")
	var replacement replacementApi

	fmt.Println("Using API lib:", viper.GetString("api_library"))

	if viper.GetString("api_library") == "Fiber" {
		replacement = FiberReplacement
	} else {
		replacement = GinReplacement
	}

	setupFile = strings.ReplaceAll(setupFile, "{{PACKAGE_NAME}}", packageName)
	setupFile = strings.ReplaceAll(setupFile, "{{IMPORT_API_LIB}}", replacement.ImportApiLib)
	setupFile = strings.ReplaceAll(setupFile, "{{IMPORT_API_LIB_MIDDLEWARE_CORS}}", replacement.ImportApiLibMiddlewareCors)
	setupFile = strings.ReplaceAll(setupFile, "{{IMPORT_API_LIB_MIDDLEWARE_LOGGER}}", replacement.ImportApiLibMiddlewareLogger)
	setupFile = strings.ReplaceAll(setupFile, "{{IMPORT_API_LIB_MIDDLEWARE_RECOVER}}", replacement.ImportApiLibRecover)
	setupFile = strings.ReplaceAll(setupFile, "{{APP_CREATION}}", replacement.AppCreation)
	setupFile = strings.ReplaceAll(setupFile, "{{API_LOGGER}}", replacement.ApiLogger)
	setupFile = strings.ReplaceAll(setupFile, "{{API_CORS}}", replacement.ApiCors)
	setupFile = strings.ReplaceAll(setupFile, "{{API_LISTEN}}", replacement.ApiListen)

	return setupFile
}

func initAppSetup() {
	filePath := "./app/setup.go"

	if _, err := os.Stat(filePath); err == nil {
		return
	}

	os.MkdirAll("./app", os.ModePerm)

	file, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	_, err = file.WriteString(writeSetupFile())
	if err != nil {
		panic(err)
	}
}
