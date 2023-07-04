package projectinit

import "github.com/ZiplEix/pear-cli/prompt"

func AskIfDocker() {
	promptContent := prompt.PromptContent{
		ErrorMsg: "Please enter a valid option.",
		Label:    "Do you want to use Docker ?",
		Items:    []string{"Yes", "No"},
	}
	usingDocker := prompt.GetSelect(promptContent, false)

	if usingDocker == "Yes" {
		SetUsingDocker(true)
		AskForGoVersion()
	} else {
		SetUsingDocker(false)
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

	SetGoVersion(goVersion)
}
