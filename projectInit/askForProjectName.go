package projectinit

import (
	"errors"
	"strings"

	"github.com/ZiplEix/pear-cli/prompt"
)

func AskForProjectName() {
	promptContent := prompt.PromptContent{
		ErrorMsg: "Please enter a valid project name.",
		Label:    "What is the name of your project ?",
		Items:    []string{},
	}
	projectName := prompt.GetInput(promptContent, func(input string) error {
		if input == "" {
			return errors.New("Project name can't be empty")
		}
		if len(input) > 50 {
			return errors.New("Project name can't be longer than 50 characters")
		}
		if strings.Contains(input, " ") {
			return errors.New("Project name can't contain spaces")
		}
		return nil
	})

	if projectName == "" {
		return
	}

	SetName(projectName)
}
