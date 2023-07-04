package projectinit

import (
	"fmt"

	"github.com/spf13/viper"
)

func PearlInit() {
	fmt.Println("pearl init called")

	initSettings()

	// ask for project name
	AskForProjectName()

	// Ask which API library to use
	AskWhichApiLibrary()

	// Ask if docker is needed in the project
	AskIfDocker()

	// Ask if a database is needed in the project
	AskIfDatabase()

	PrintSettings()

	viper.WriteConfig()
}
