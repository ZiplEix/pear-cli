/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	projectinit "github.com/ZiplEix/pear-cli/projectInit"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize your new API project.",
	Long:  `Initialize your new API project.`,
	Run: func(cmd *cobra.Command, args []string) {
		projectinit.PearlInit()
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// intiCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// intiCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
