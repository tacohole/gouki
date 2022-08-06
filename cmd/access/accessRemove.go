package access

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	herokuApi "github.com/tacohole/gouki/util/heroku-api"
)

var AccessRemoveCmd = &cobra.Command{
	Short: "remove users from a team app",
	Long: `USAGE
  $ gouki access remove -a APP -u USER`,
	Use:  "remove",
	Args: cobra.MaximumNArgs(1),
	Run:  accessRemove,
}

func init() {
	AccessRemoveCmd.Flags().StringVarP(&user, "user", "u", "", "user email to be removed from the app")
	AccessRemoveCmd.MarkFlagRequired("user")
	AccessRemoveCmd.MarkFlagRequired("app")

	AccessCmd.AddCommand(AccessRemoveCmd)
}

func accessRemove(cmd *cobra.Command, args []string) {
	app := viper.GetString("app")

	client := herokuApi.MakeClient()

	fmt.Printf("Removing %s access from the app %s...", user, app)

	// TODO: handle team collab removal, it's probably different
	if err := client.CollaboratorDelete(app, user); err != nil {
		log.Fatalf("%v", err)
	}
	fmt.Print("done")
}
