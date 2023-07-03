package projectinit

import "fmt"

type PearlSettings struct {
	Name        string
	Replacement map[string]string
	Docker      bool
}

var Settings PearlSettings

func (s *PearlSettings) initSettings() {
	s.Name = "Go-API-with-pearl"
	s.Replacement = map[string]string{
		"{{GO_VERSION}}": "1.20.2",
	}
	s.Docker = true
}

func (s *PearlSettings) PrintSettings() {
	fmt.Println("Settings = {")
	fmt.Println("  Name:", s.Name)
	fmt.Println("  Replacement:", s.Replacement)
	fmt.Println("  Docker:", s.Docker)
	fmt.Println("}")
}

func (s *PearlSettings) ModifReplacement(key, value string) {
	s.Replacement[key] = value
}

func (s *PearlSettings) setDocker(docker bool) {
	s.Docker = docker
}

func (s *PearlSettings) SetName(name string) {
	s.Name = name
}
