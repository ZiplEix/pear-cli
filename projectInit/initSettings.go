package projectinit

import "fmt"

type PearlSettings struct {
	Name          string
	Replacement   map[string]string
	ApiLibrary    string
	UsingDocker   bool
	UsingDatabase bool
	Database      string
	Orm           string
}

var Settings PearlSettings

func (s *PearlSettings) initSettings() {
	s.Name = "Go-API-with-pearl"
	s.Replacement = map[string]string{
		"{{GO_VERSION}}": "1.20.2",
	}
	s.ApiLibrary = "Fiber"
	s.UsingDocker = true
	s.UsingDatabase = false
	s.Database = ""
	s.Orm = ""
}

func (s *PearlSettings) PrintSettings() {
	fmt.Println("Settings = {")
	fmt.Println("\tName:", s.Name)
	fmt.Println("\tReplacement:", s.Replacement)
	fmt.Println("\tapiLibrary:", s.ApiLibrary)
	fmt.Println("\tUsingDocker:", s.UsingDocker)
	fmt.Println("\tUsingDatabase:", s.UsingDatabase)
	fmt.Println("\tDatabase:", s.Database)
	fmt.Println("\tOrm:", s.Orm)
	fmt.Println("}")
}

func (s *PearlSettings) setApiLibrary(apiLibrary string) {
	s.ApiLibrary = apiLibrary
}

func (s *PearlSettings) SetName(name string) {
	s.Name = name
}

func (s *PearlSettings) ModifReplacement(key, value string) {
	s.Replacement[key] = value
}

func (s *PearlSettings) setUsingDocker(docker bool) {
	s.UsingDocker = docker
}

func (s *PearlSettings) setUsingDatabase(database bool) {
	s.UsingDatabase = database
}

func (s *PearlSettings) setDatabase(database string) {
	s.Database = database
}

func (s *PearlSettings) setOrm(orm string) {
	s.Orm = orm
}
