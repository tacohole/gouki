package access

import "github.com/spf13/cobra"

var AccessAddCmd = &cobra.Command{
	Short: "add new users to your app",
	Long: `USAGE
  $ gouki access:add EMAIL

OPTIONS
  -a, --app=app                  (required) app to run command against
  -p, --permissions=permissions  list of permissions comma separated
  -r, --remote=remote            git remote of app to use

EXAMPLES
  $ gouki access add user@email.com --app APP # add a collaborator to your app
  $ gouki access add user@email.com --app APP --permissions
  deploy,manage,operate # permissions must be comma separated`,
	Use: "add",
	Run: accessAdd,
}

var permissions []string

func init() {
	AccessAddCmd.Flags().StringArrayVarP(&permissions, "permissions", "p", []string{}, "")

	AccessCmd.AddCommand(AccessAddCmd)

}

func accessAdd(cmd *cobra.Command, args []string) {}
