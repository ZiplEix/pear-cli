package project

import (
	"fmt"
	"os"
	"text/template"
)

type EnvData struct {
}

func newEnvData() *AirData {
	return &AirData{}
}

func (p *Project) initEnv() error {
	data := newEnvData()

	tmpl, err := template.ParseFiles("../templates/env.tmpl")
	if err != nil {
		return fmt.Errorf("Error parsing template: %v", err)
	}

	file, err := os.Create(".env")
	if err != nil {
		return fmt.Errorf("Error opening env file: %v", err)
	}
	defer file.Close()

	err = tmpl.Execute(file, data)
	if err != nil {
		return fmt.Errorf("Error executing template: %v", err)
	}

	fmt.Println("Env file created")

	return nil
}
