package project

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/sanity-io/litter"
	"gopkg.in/yaml.v3"
)

type Project struct {
	Name         string `yaml:"name"`
	Path         string `yaml:"path"`
	Docker       bool   `yaml:"useDockerfile"`
	Compose      bool   `yaml:"useDockerCompose"`
	Air          bool   `yaml:"useAir"`
	Swagger      bool   `yaml:"useSwagger"`
	ApiFramework string `yaml:"apiFramework"`
}

func NewProject(name string, path string, docker, air, swagger, compose bool, framework string) *Project {
	return &Project{
		Name:         name,
		Path:         path,
		Docker:       docker,
		Compose:      compose,
		Air:          air,
		Swagger:      swagger,
		ApiFramework: framework,
	}
}

func (p Project) toYaml() error {
	data, err := yaml.Marshal(&p)
	if err != nil {
		return fmt.Errorf("Error formating yaml: %v", err)
	}

	file, err := os.OpenFile(".peach.yaml", os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return fmt.Errorf("Error opening file: %v", err)
	}
	defer file.Close()

	_, err = file.Write(data)
	if err != nil {
		return fmt.Errorf("Error writing to file: %v", err)
	}

	fmt.Println("Config file created")

	return nil
}

func LoadFromYaml() (*Project, error) {
	var project Project

	file, err := os.Open(".peach.yaml")
	if err != nil {
		return nil, fmt.Errorf("Can't open config file: %v", err)
	}

	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&project)
	if err != nil {
		return nil, fmt.Errorf("Can't decode yaml file: %v", err)
	}

	return &project, nil
}

func (p *Project) createProjectDirectory(force bool) error {
	if _, err := os.Stat(p.Path); os.IsNotExist(err) {
		err := os.Mkdir(p.Path, 0755)
		if err != nil {
			panic(err)
		}
	} else {
		var response string
		if !force {
			fmt.Printf("Directory %s already exist. Do you want to erase it ? (y/n): ", p.Path)
			reader := bufio.NewReader(os.Stdin)
			response, _ = reader.ReadString('\n')
		}
		if strings.TrimSpace(response) == "y" || force {
			err := os.RemoveAll(p.Path)
			if err != nil {
				panic(err)
			}
			err = os.Mkdir(p.Path, 0755)
			if err != nil {
				panic(err)
			}
		} else {
			fmt.Println("Aborted")
			os.Exit(0)
		}
	}

	return nil
}

func (p *Project) Init(force bool) {
	litter.Dump(p)

	if err := p.createProjectDirectory(force); err != nil {
		panic(err)
	}

	if err := os.Chdir(p.Path); err != nil {
		panic(err)
	}

	cmd := exec.Command("go", "mod", "init", p.Name)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		panic(err)
	}

	if err := p.toYaml(); err != nil {
		panic(err)
	}

	if p.Docker {
		if err := p.initDocker(); err != nil {
			panic(err)
		}
	}

	if p.Compose {
		if err := p.initCompose(); err != nil {
			panic(err)
		}
	}

	if p.Air {
		if err := p.initAir(); err != nil {
			panic(err)
		}
	}

	if err := p.initEnv(); err != nil {
		panic(err)
	}

	if err := p.initCodeFile(); err != nil {
		panic(err)
	}

	if p.Swagger {
		if err := generateFirstDoc(); err != nil {
			panic(err)
		}
	}

	if err := downloadPackages(); err != nil {
		panic(err)
	}

	if err := cleanGeneratedGoFiles(); err != nil {
		panic(err)
	}
}
