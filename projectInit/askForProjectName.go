package projectinit

import "github.com/ZiplEix/pear-cli/prompt"

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
