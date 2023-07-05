package initfiles

import (
	"io/ioutil"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/viper"
)

func initGoProject() {
	// if the file ./go.mod exists, we don't need to run the command "go mod init"
	if _, err := os.Stat("./go.mod"); err == nil {
		return
	}

	cmd := exec.Command("go", "mod", "init")
	err := cmd.Run()
	if err != nil {
		cmd = exec.Command("go", "mod", "init", viper.GetString("name"))
		err = cmd.Run()
		if err != nil {
			panic(err)
		}
	}

	// get the name of the package
	goMod, err := ioutil.ReadFile("./go.mod")
	if err != nil {
		panic(err)
	}

	packageName := strings.Split(strings.Split(string(goMod), "\n")[0], " ")[1]

	viper.Set("package_name", packageName)
}

func Init() {
	initGoProject()

	if viper.GetBool("using_docker") {
		initDockerFile()
	}

	initAirConfigFile()
}
