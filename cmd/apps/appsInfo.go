package apps

import "github.com/spf13/cobra"

var appsInfoCmd = &cobra.Command{
	Short: "retrieves details information on a Heroku app",
	Long:  "",
	Use:   "info",
	Run:   getAppInfo,
}

// OPTIONS
//   -a, --app=app          app to run command against
//   -c, --confirm=confirm
//   -r, --remote=remote    git remote of app to use

func init() {
	AppsCmd.AddCommand(appsInfoCmd)
}

func getAppInfo(cmd *cobra.Command, args []string) {
	loadDefaultVariables()
}
