package project

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"
)

var (
	files = []string{
		"main",
		"routes/main",
		"routes/version",
	}
)

type CodeFileData struct {
	Swagger            bool
	ProjectName        string
	ProjectDescription string
	Framework          string
}

func newCodeFileData(swagger bool, projectName string, framework string) *CodeFileData {
	return &CodeFileData{
		Swagger:            swagger,
		ProjectName:        projectName,
		ProjectDescription: "Description of your project for the swagger doc",
		Framework:          framework,
	}
}

func generateCodeFile(data *CodeFileData, fileName, tmplFile string) error {
	tmpl, err := template.ParseFiles(fmt.Sprintf("../templates/%s/%s", data.Framework, tmplFile))
	if err != nil {
		return fmt.Errorf("Error parsing template: %v", err)
	}

	dir := filepath.Dir(fileName)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return fmt.Errorf("Error: fail to create folder for %s: %v", fileName, err)
	}

	file, err := os.Create(fileName)
	if err != nil {
		return fmt.Errorf("Error opening %s config file: %v", fileName, err)
	}
	defer file.Close()

	err = tmpl.Execute(file, data)
	if err != nil {
		return fmt.Errorf("Error executing template: %v", err)
	}

	return nil
}

func (p *Project) initCodeFile() error {
	data := newCodeFileData(p.Swagger, p.Name, p.ApiFramework)

	for _, file := range files {
		fileName := file + ".go"
		tmplFile := file + ".tmpl"
		if err := generateCodeFile(data, fileName, tmplFile); err != nil {
			return fmt.Errorf("Error generating %s code file: %v", fileName, err)
		}
		fmt.Println(fileName + " created")
	}

	return nil
}
