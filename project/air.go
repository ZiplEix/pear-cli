package project

import (
	"fmt"
	"os"
	"text/template"
)

type AirData struct {
	Swagger bool
}

func newAirData(swagger bool) *AirData {
	return &AirData{
		Swagger: swagger,
	}
}

func (p *Project) initAir() error {
	data := newAirData(p.Swagger)

	tmpl, err := template.ParseFiles("../templates/air.tmpl")
	if err != nil {
		return fmt.Errorf("Error parsing template: %v", err)
	}

	file, err := os.Create(".air.toml")
	if err != nil {
		return fmt.Errorf("Error opening air config file: %v", err)
	}
	defer file.Close()

	err = tmpl.Execute(file, data)
	if err != nil {
		return fmt.Errorf("Error executing template: %v", err)
	}

	return nil
}
