package apps

import "github.com/spf13/cobra"

var createAppCmd = &cobra.Command{
	Short: "creates a new Heroku app",
	Long:  "creates a new Heroku app",
	Use:   "create",
	Run:   createApp,
}

// OPTIONS
//   -b, --buildpack=buildpack  buildpack url to use for this app
//   -n, --no-remote            do not create a git remote
//   -r, --remote=remote        the git remote to create, default "heroku"
//   -s, --stack=stack          the stack to create the app on
//   -t, --team=team            team to use
//   --addons=addons            comma-delimited list of addons to install
//   --json                     output in json format
//   --region=region            specify region for the app to run in
//   --space=space              the private space to create the app in

func init() {
	AppsCmd.AddCommand(createAppCmd)
}

func createApp(cmd *cobra.Command, args []string) {
	loadDefaultVariables()
}
