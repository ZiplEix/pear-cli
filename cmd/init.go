package cmd

import (
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
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Init a new classic go api project",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		project := project.NewProject(name, path, docker, air, swagger, dockerCompose)
		project.Init(force)
	},
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

	initCmd.MarkFlagRequired("name")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this commanJJd
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
