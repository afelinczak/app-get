package cmd

import (
	"fmt"
	"strings"

	"github.com/afelinczak/app-get/domain"
	"github.com/afelinczak/app-get/infrastructure"
	"github.com/spf13/cobra"
)

// installCmd represents the install command
var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install program defined in repository",
	Long:  `Use search or list command to see list of available apps.`,
	Run: func(cmd *cobra.Command, args []string) {
		if strings.Index(args[0], "/") == -1 {
			fmt.Println("To install app from github you need to pass userName/repoName")
			return
		}

		fmt.Println("Check if deb is available")
		var app = domain.App{SourceUrl: args[0], Name: strings.Split(args[0], "/")[1], AppType: domain.Deb, Source: domain.Github}
		var version = infrastructure.GetLatestVersion(app)
		var path = infrastructure.GetInstallationFile(app, version)
		infrastructure.InstallApp(path)
		domain.AddNewApp(app, version.Name, infrastructure.GetInstalledApps, infrastructure.WriteInstalledApps)
		fmt.Println(app.Name + " installed")
	},
}

func init() {
	rootCmd.AddCommand(installCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// installCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// installCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
