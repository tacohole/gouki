package access

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	herokuApi "github.com/tacohole/gouki/util/heroku-api"
)

var AccessUpdateCmd = &cobra.Command{
	Short: "update existing collaborators on an team app",
	Long: `USAGE
  $ gouki access update -u EMAIL -a APP -p PERMISSIONS`,
	Use: "update",
	Run: accessUpdate,
}

func init() {
	AccessAddCmd.Flags().StringArrayVarP(&permissions, "permissions", "p", []string{}, "list of permissions comma separated")
	AccessAddCmd.Flags().StringVarP(&user, "user", "u", "", "email of collaborator")

	AccessAddCmd.MarkFlagRequired("user")
	AccessAddCmd.MarkFlagRequired("app")
	AccessCmd.AddCommand(AccessUpdateCmd)
}

func accessUpdate(cmd *cobra.Command, args []string) {
	app := viper.GetString("app")

	_, err := herokuApi.UpdateTeamCollaborator(app, user, permissions)
	if err != nil {
		log.Fatalf("%v", err)
	}
}
