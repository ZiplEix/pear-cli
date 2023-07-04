package projectinit

import "github.com/ZiplEix/pear-cli/prompt"

func selectDatabase() {
	promptContent := prompt.PromptContent{
		ErrorMsg: "Please enter a valid option.",
		Label:    "Select your database",
		Items:    []string{"PostgreSQL", "MySQL", "SqLite"},
	}
	database := prompt.GetSelect(promptContent, false)

	SetDatabase(database)
}

func selectOrm() {
	promptContent := prompt.PromptContent{
		ErrorMsg: "Please enter a valid option.",
		Label:    "Select your ORM",
		Items:    []string{"Gorm", "Xorm"},
	}
	orm := prompt.GetSelect(promptContent, false)

	SetOrm(orm)
}

func AskIfDatabase() {
	promptContent := prompt.PromptContent{
		ErrorMsg: "Please enter a valid option.",
		Label:    "Do you want to use a database ?",
		Items:    []string{"Yes", "No"},
	}
	usingDatabase := prompt.GetSelect(promptContent, false)

	if usingDatabase == "Yes" {
		SetUsingDatabase(true)
		selectDatabase()
		selectOrm()
	} else {
		SetUsingDatabase(false)
	}
}
