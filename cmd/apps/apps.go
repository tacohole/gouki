package apps

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tacohole/gouki/util/heroku"
)

var AppsCmd = &cobra.Command{
	Use:   "apps",
	Short: "get info about all apps",
	Long:  "get information about all the Heroku apps you own",
	Args:  cobra.MaximumNArgs(1),
	Run:   apps,
}

// OPTIONS
//   -A, --all          include apps in all teams
//   -p, --personal     list apps in personal account when a default team is set
//   -s, --space=space  filter by space
//   -t, --team=team    team to use
//   --json             output in json format

// var verbose bool
var allApps string
var personalApps string
var space string
var team string

func init() {
	AppsCmd.Flags().StringVarP(&allApps, "all", "a", "", "include apps in all teams")
	AppsCmd.Flags().StringVarP(&personalApps, "personal", "p", "", "list apps in personal account when a default team is set")
	AppsCmd.Flags().StringVarP(&space, "space", "s", "", "filter by space")
	AppsCmd.Flags().StringVarP(&team, "team", "t", "", "team to use")

}

func apps(cmd *cobra.Command, args []string) {
	loadDefaultVariables()
	// setup client
	client, err := heroku.MakeClient()
	if err != nil {
		fmt.Printf("error making client: %s", err)
	}
	// get apps
	resp, err := client.AppList(nil)
	if err != nil {
		fmt.Printf("error retrieving apps: %s", err)
	}

	// print that ish
	fmt.Printf("=== %s Apps\n", client.Username)
	for _, app := range resp {
		fmt.Printf("%s \n", app.Name)

	}
	fmt.Printf("\n")

}

func loadDefaultVariables() {
	// verbose = viper.GetBool("VERBOSE")
}
