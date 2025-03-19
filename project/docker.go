package project

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"text/template"
)

type DockerFileData struct {
	BaseImage string
	Swagger   bool
}

func newDockerFileData(swagger bool) *DockerFileData {
	return &DockerFileData{
		BaseImage: "golang:1.24",
		Swagger:   swagger,
	}
}

type DockerComposeData struct {
	Redis    bool
	Postgres bool
}

func newComposeData(redis, postgres bool) *DockerComposeData {
	return &DockerComposeData{
		Redis:    redis,
		Postgres: postgres,
	}
}

func (p *Project) initDocker() error {
	data := newDockerFileData(p.Swagger)

	tmpl, err := template.ParseFiles("../templates/Dockerfile.tmpl")
	if err != nil {
		return fmt.Errorf("Error parsing template: %v", err)
	}

	// get go version
	cmd := exec.Command("go", "version")
	out, err := cmd.Output()
	if err != nil {
		return fmt.Errorf("Error getting go version: %v", err)
	}

	versionRe := regexp.MustCompile(`go(\d+\.\d+(?:\.\d+)?)`)
	version := versionRe.FindStringSubmatch(string(out))[1]

	data.BaseImage = fmt.Sprintf("golang:%s", version)

	file, err := os.Create("Dockerfile")
	if err != nil {
		return fmt.Errorf("Error opening Dockerfile: %v", err)
	}
	defer file.Close()

	err = tmpl.Execute(file, data)
	if err != nil {
		return fmt.Errorf("Error executing template: %v", err)
	}

	fmt.Println("Dockerfile created")

	return nil
}

func (p *Project) initCompose() error {
	data := newComposeData(false, false)

	tmpl, err := template.ParseFiles("../templates/docker-compose.tmpl")
	if err != nil {
		return fmt.Errorf("Error parsing template: %v", err)
	}

	file, err := os.Create("docker-compose.yml")
	if err != nil {
		return fmt.Errorf("Error opening docker-compose.yml: %v", err)
	}
	defer file.Close()

	err = tmpl.Execute(file, data)
	if err != nil {
		return fmt.Errorf("Error executing template: %v", err)
	}

	fmt.Println("docker-compose.yml created")

	return nil
}
