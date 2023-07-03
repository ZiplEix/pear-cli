package prompt

import (
	"fmt"

	"github.com/manifoldco/promptui"
)

func GetSelect(pc PromptContent, other bool) string {
	index := -1

	var result string
	var err error

	for index < 0 {
		if other == false {
			prompt := promptui.Select{
				Label: pc.Label,
				Items: pc.Items,
			}
			index, result, err = prompt.Run()
		} else {
			prompt := promptui.SelectWithAdd{
				Label:    pc.Label,
				Items:    pc.Items,
				AddLabel: "Other",
			}

			index, result, err = prompt.Run()

			if index == -1 {
				break
			}
		}
	}

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return ""
	}

	return result
}
