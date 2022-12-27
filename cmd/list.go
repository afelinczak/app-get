/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/afelinczak/app-get/infrastructure"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Display list of installed apps",
	Run: func(cmd *cobra.Command, args []string) {
		var appRepo = infrastructure.AppRepository{}
		var apps = appRepo.Get()
		fmt.Println("Installed applications")
		for i := 0; i < len(apps); i++ {
			fmt.Println(apps[i].App.Name + " (" + apps[i].Version + ")")
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
