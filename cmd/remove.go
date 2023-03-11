/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"strings"

	"github.com/afelinczak/app-get/domain"
	"github.com/afelinczak/app-get/infrastructure"
	"github.com/spf13/cobra"
)

// removeCmd represents the remove command
var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Removes repository from list of apps to update",
	Long:  `Use sudo app-get remove githubUser/githubRepository to remove app update list. Please note it will not unistall the app`,
	Run: func(cmd *cobra.Command, args []string) {
		var appRepo = infrastructure.AppRepository{}
		if len(args) == 0 || strings.Index(args[0], "/") == -1 {
			fmt.Println("To remove app from app-get you need to pass userName/repoName")
			return
		}

		domain.RemoveApp(args[0], appRepo)
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// removeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// removeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
