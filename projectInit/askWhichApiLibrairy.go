package projectinit

import "github.com/ZiplEix/pear-cli/prompt"

func AskWhichApiLibrary() {
	promptContent := prompt.PromptContent{
		ErrorMsg: "Please enter a valid option.",
		Label:    "Select your API library",
		Items:    []string{"Fiber", "Gin", "Echo"},
	}
	apiLibrary := prompt.GetSelect(promptContent, false)

	Settings.setApiLibrary(apiLibrary)
}
