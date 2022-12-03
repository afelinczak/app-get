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
	Short: "refresh list of applications from repository",
	Run: func(cmd *cobra.Command, args []string) {
		var installedApps = infrastructure.GetInstalledApps()
		for i := 0; i < len(installedApps); i++ {
			var version = infrastructure.GetLatestVersion(installedApps[i].App)
			fmt.Println("Latest available version of " + installedApps[i].App.Name + " version is " + version.Name)

			if domain.IsAvailableVersionNewerAndStable(installedApps[i].Version, version.Name) {
				fmt.Println("Download newer version of " + installedApps[i].App.Name)

				var path = infrastructure.GetInstallationFile(installedApps[i].App, version)
				infrastructure.InstallApp(path)
			}
		}
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
