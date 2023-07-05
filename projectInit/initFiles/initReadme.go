package initfiles

import (
	"os"

	"github.com/spf13/viper"
)

func initReadme() {
	if _, err := os.Stat("./README.md"); err == nil {
		return
	}

	file, err := os.Create("./README.md")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	_, err = file.WriteString("# " + viper.GetString("name") + "\n")
	if err != nil {
		panic(err)
	}
}
