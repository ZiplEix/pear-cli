package initfiles

import (
	"os"
	"strings"

	"github.com/spf13/viper"
)

var dockerfileTemplate = "FROM golang:{{GO_VERSION}} as builder\n\nWORKDIR /usr/src/app\n\nRUN go install github.com/cosmtrek/air@latest\nRUN go install github.com/swaggo/swag/cmd/swag@latest\n\nCOPY . .\n\nRUN go mod tidy\n"

func writeDockerfile() string {
	dockerfile := strings.ReplaceAll(dockerfileTemplate, "{{GO_VERSION}}", viper.GetString("replacement.go_version"))

	return dockerfile
}

func initDockerFile() {
	// if the file ./Dockerfile exists, we don't need to create it
	if _, err := os.Stat("./Dockerfile"); err == nil {
		return
	}

	file, err := os.Create("./Dockerfile")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	_, err = file.WriteString(writeDockerfile())
	if err != nil {
		panic(err)
	}
}
