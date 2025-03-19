package cmd

import (
	"fmt"

	"github.com/ZiplEix/pear_cli/project"
	"github.com/spf13/cobra"
)

var (
	name          string
	path          string
	force         bool
	air           bool
	docker        bool
	dockerCompose bool
	swagger       bool
	framework     string
	full          bool
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Init a new classic go api project",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if needRepl(cmd) && !full {
			repl()
		}

		if full {
			setAllFeatures()
		}

		project := project.NewProject(name, path, docker, air, swagger, dockerCompose, framework)
		project.Init(force)
	},
}

func setAllFeatures() {
	path = "./" + name
	force = false
	air = true
	docker = true
	dockerCompose = true
	swagger = true
	framework = "fiber"
}

func needRepl(cmd *cobra.Command) bool {
	return !cmd.Flags().Changed("name") ||
		(!cmd.Flags().Changed("path") &&
			!cmd.Flags().Changed("air") &&
			!cmd.Flags().Changed("docker") &&
			!cmd.Flags().Changed("docker-compose") &&
			!cmd.Flags().Changed("swagger"))
}

func ynToBool(yn string) bool {
	return yn == "y" || yn == "yes"
}

func repl() {
	// Ask for project name and path
	fmt.Print("Enter project name: ")
	_, _ = fmt.Scanln(&name)

	fmt.Print("Enter project path (default './'): ")
	_, _ = fmt.Scanln(&path)
	if path == "" || path == "\n" {
		path = "./"
	}

	// Ask for project features
	var res string

	fmt.Print("Do you want to use air for daemonize the server? (y/n): ")
	_, _ = fmt.Scanln(&res)
	air = ynToBool(res)

	fmt.Print("Do you want to use docker for containerize the server? (y/n): ")
	_, _ = fmt.Scanln(&res)
	docker = ynToBool(res)

	fmt.Print("Do you want to use docker-compose for containerize the server and his dependencies? (y/n): ")
	_, _ = fmt.Scanln(&res)
	dockerCompose = ynToBool(res)

	fmt.Print("Do you want to use swagger for documentation? (y/n): ")
	_, _ = fmt.Scanln(&res)
	swagger = ynToBool(res)

	// Ask for project framework
	fmt.Print("Enter project framework (default 'fiber'): ")
	_, _ = fmt.Scanln(&framework)
	if framework == "" || framework == "\n" {
		framework = "fiber"
	}

	fmt.Println()
}

func init() {
	rootCmd.AddCommand(initCmd)

	initCmd.Flags().StringVarP(&name, "name", "n", "", "Project name (will be appended to the command 'go mod init')")
	initCmd.Flags().StringVarP(&path, "path", "p", "./", "Project path")
	initCmd.Flags().BoolVarP(&force, "force", "f", false, "Force the creation of the project even if the directory already exist")
	initCmd.Flags().BoolVar(&air, "air", false, "Use air for daemonize the server")
	initCmd.Flags().BoolVar(&docker, "docker", false, "Use docker for containerize the server")
	initCmd.Flags().BoolVar(&dockerCompose, "docker-compose", false, "Use docker-compose for containerize the server and his dependencies")
	initCmd.Flags().BoolVar(&swagger, "swagger", false, "Use swagger for documentation")
	initCmd.Flags().StringVar(&framework, "framework", "fiber", "Use a specific framework for the project")
	initCmd.Flags().BoolVar(&full, "full", false, "Create a full project with all the features")
}
