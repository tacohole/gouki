package access

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	herokuApi "github.com/tacohole/gouki/util/heroku-api"
)

var AccessAddCmd = &cobra.Command{
	Short: "add new users to your app",
	Long: `USAGE
  $ gouki access add EMAIL

EXAMPLES
  $ gouki access add -u user@email.com --a APP # add a collaborator to your app
  $ gouki access add -u user@email.com --a APP --permissions deploy,manage,operate 
	# permissions must be comma separated`,
	Use: "add",
	Run: accessAdd,
}

var permissions []string
var silent bool

func init() {
	AccessAddCmd.Flags().StringArrayVarP(&permissions, "permissions", "p", []string{}, "list of permissions comma separated")
	AccessAddCmd.Flags().StringVarP(&user, "user", "u", "", "email of new collaborator")
	AccessAddCmd.Flags().BoolVarP(&silent, "silent", "s", false, "suppress email invitation to collaborator")

	AccessAddCmd.MarkFlagRequired("user")
	AccessAddCmd.MarkFlagRequired("app")

	AccessCmd.AddCommand(AccessAddCmd)

}

func accessAdd(cmd *cobra.Command, args []string) {
	client := herokuApi.MakeClient()

	app := viper.GetString("app")

	appInfo, err := client.AppInfo(app)
	if err != nil {
		log.Fatalf("%v", err)
	}

	fmt.Printf("Adding %s access to the app %s...", user, app)

	if herokuApi.IsTeamApp(appInfo.Owner.Email) {
		_, err := herokuApi.AddTeamCollaborator(app, user, permissions, silent)
		if err != nil {
			log.Fatalf("%v", err)
		}
		fmt.Println("done")
	}

	opts := herokuApi.MakeCollaboratorOpts(silent)

	_, err = client.CollaboratorCreate(app, user, opts)
	if err != nil {
		log.Printf("%s", err)
	}
	fmt.Println("done")
}
