package projectinit

import (
	"fmt"

	"github.com/spf13/viper"
)

func initSettings() {
	viper.SetConfigName(".pear")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("Config file not found")
		} else {
			fmt.Println("Config file found but another error occurred")
		}
	}

	viper.SetDefault("name", "Go-API-with-pearl")
	viper.SetDefault("api_library", "Fiber")
	viper.SetDefault("using_docker", true)
	viper.SetDefault("replacement.go_version", "1.20.2")
	viper.SetDefault("using_database", false)
	viper.SetDefault("database", "")
	viper.SetDefault("orm", "")
}

func PrintSettings() {
	// fmt.Println("Settings = {")
	// fmt.Println("    Name:", viper.GetString("name"))
	// fmt.Println("    Replacement:", viper.GetStringMapString("replacement"))
	// fmt.Println("    GO_VERSION:", viper.GetString("replacement.GO_VERSION"))
	// fmt.Println("    apiLibrary:", viper.GetString("apiLibrary"))
	// fmt.Println("    UsingDocker:", viper.GetBool("usingDocker"))
	// fmt.Println("    UsingDatabase:", viper.GetBool("usingDatabase"))
	// fmt.Println("    Database:", viper.GetString("database"))
	// fmt.Println("    Orm:", viper.GetString("orm"))
	// fmt.Println("}")
}

func SetApiLibrary(apiLibrary string) {
	viper.Set("api_library", apiLibrary)
}

func SetName(name string) {
	viper.Set("name", name)
}

func SetUsingDocker(docker bool) {
	viper.Set("using_docker", docker)
}

func SetGoVersion(version string) {
	viper.Set("replacement.go_version", version)
}

func SetUsingDatabase(database bool) {
	viper.Set("using_database", database)
}

func SetDatabase(database string) {
	viper.Set("database", database)
}

func SetOrm(orm string) {
	viper.Set("orm", orm)
}
