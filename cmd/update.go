package cmd

import (
	"fmt"

	"github.com/afelinczak/app-get/domain"
	"github.com/afelinczak/app-get/infrastructure"
	"github.com/spf13/cobra"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "download new version of the installed apps and installs them",
	Run: func(cmd *cobra.Command, args []string) {
		var appRepo = infrastructure.AppRepository{}
		var installedApps = appRepo.Get()

		for i := 0; i < len(installedApps); i++ {
			var version = infrastructure.GetLatestVersion(installedApps[i].App)

			if domain.IsAvailableVersionNewerAndStable(installedApps[i].Version, version.Name) {
				fmt.Println("Download " + installedApps[i].App.Name + " " + version.Name)

				var path, success = infrastructure.GetInstallationFile(installedApps[i].App, version)
				if success {
					infrastructure.InstallApp(path)
					domain.UpdateAppVersion(installedApps[i].App, version.Name, appRepo)
				}
			} else {
				fmt.Println("No newer version of " + installedApps[i].App.Name + " found (" + version.Name + ")")
			}
		}
		fmt.Println("Update finished")
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// updateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// updateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
