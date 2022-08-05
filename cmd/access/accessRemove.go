package access

import "github.com/spf13/cobra"

var AccessRemoveCmd = &cobra.Command{
	Short: "list who has access to an app",
	Long: `USAGE
  $ gouki access

OPTIONS
  -a, --app=app        (required) app to run command against
  -r, --remote=remote  git remote of app to use
  --json               output in json format`,
	Use: "access:remove",
	Run: accessRemove,
}

func init() {

}

func accessRemove(cmd *cobra.Command, args []string) {}
