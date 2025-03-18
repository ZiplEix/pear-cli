package project

import (
	"fmt"
	"os"
	"os/exec"
)

func downloadPackages() error {
	cmd := exec.Command("go", "mod", "tidy")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("Error executing 'go mod tidy' command: %v", err)
	}

	return nil
}

func cleanGeneratedGoFiles() error {
	cmd := exec.Command("go", "fmt")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("Error executing 'go fmt' command: %v", err)
	}

	return nil
}

func generateFirstDoc() error {
	installSwagger := exec.Command("go", "install", "github.com/go-swagger/go-swagger/cmd/swagger@latest")
	installSwagger.Stdout = os.Stdout
	installSwagger.Stderr = os.Stderr

	err := installSwagger.Run()
	if err != nil {
		return fmt.Errorf("Error installing swagger: %v", err)
	}

	cmd := exec.Command("swag", "init", "--parseDependency", "--parseInternal")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("Error generating swagger doc: %v", err)
	}

	return nil
}
