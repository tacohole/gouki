package access

import "github.com/spf13/cobra"

var AccessUpdateCmd = &cobra.Command{
	Short: "list who has access to an app",
	Long: `USAGE
  $ gouki access

OPTIONS
  -a, --app=app        (required) app to run command against
  -r, --remote=remote  git remote of app to use
  --json               output in json format`,
	Use: "access:update",
	Run: accessUpdate,
}

func init() {
	AccessUpdateCmd.Flags().StringVarP(&appName, "app", "a", "", "")
	AccessUpdateCmd.Flags().StringVarP(&remote, "remote", "r", "", "")

}

func accessUpdate(cmd *cobra.Command, args []string) {}
