package apps

import "github.com/spf13/cobra"

var destroyAppCmd = &cobra.Command{
	Short: "destroys a Heroku app",
	Long:  "destroys a Heroku app",
	Use:   "destroy",
	Run:   destroyApp,
}

// OPTIONS
//   -a, --app=app          app to run command against
//   -c, --confirm=confirm
//   -r, --remote=remote    git remote of app to use

func init() {
	AppsCmd.AddCommand(destroyAppCmd)
}

func destroyApp(cmd *cobra.Command, args []string) {

}
