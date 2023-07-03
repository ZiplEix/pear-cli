package projectinit

import (
	"fmt"

	"github.com/ZiplEix/pear-cli/prompt"
)

func AskForProjectName() {
	promptContent := prompt.PromptContent{
		ErrorMsg: "Please enter a valid project name.",
		Label:    "What is the name of your project ?",
		Items:    []string{},
	}
	projectName := prompt.GetInput(promptContent, func(input string) error {
		return nil
	})

	if projectName == "" {
		return
	}

	Settings.SetName(projectName)
}

func AskIfDocker() {
	promptContent := prompt.PromptContent{
		ErrorMsg: "Please enter a valid option.",
		Label:    "Do you want to use Docker ?",
		Items:    []string{"Yes", "No"},
	}
	usingDocker := prompt.GetSelect(promptContent, false)

	if usingDocker == "Yes" {
		Settings.setDocker(true)
		AskForGoVersion()
	} else {
		Settings.setDocker(false)
	}
}

func AskForGoVersion() {
	promptContent := prompt.PromptContent{
		ErrorMsg: "Please enter a valid version.",
		Label:    "What version of Go do you want to use ?",
		Items:    []string{"1.20.2", "1.19.2", "1.18.2"},
	}
	goVersion := prompt.GetSelect(promptContent, true)

	if goVersion == "" {
		return
	}

	Settings.ModifReplacement("{{GO_VERSION}}", goVersion)
}

func PearlInit() {
	fmt.Println("pearl init called")

	Settings.initSettings()

	// ask for project name
	AskForProjectName()

	// Ask if docker is needed in the project
	AskIfDocker()

	Settings.PrintSettings()
}
