package projectinit

import (
	"fmt"

	initfiles "github.com/ZiplEix/pear-cli/projectInit/initFiles"
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

	// Print settings
	PrintSettings()

	// start the files generation
	initfiles.Init()

	// write the config in the .pear.yaml file
	viper.WriteConfigAs(".pear.yaml")
}
