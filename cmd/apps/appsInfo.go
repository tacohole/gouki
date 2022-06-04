package apps

import "github.com/spf13/cobra"

var appsInfoCmd = &cobra.Command{
	Short: "retrieves details information on a Heroku app",
	Long:  "",
	Use:   "info",
	Run:   getAppInfo,
}

// OPTIONS
//   -a, --app=app        app to run command against
//   -j, --json
//   -r, --remote=remote  git remote of app to use
//   -s, --shell          output more shell friendly key/value pairs

func init() {
	AppsCmd.AddCommand(appsInfoCmd)
}

func getAppInfo(cmd *cobra.Command, args []string) {
	loadDefaultVariables()
}
