package apps

import (
	"fmt"

	heroku "github.com/bgentry/heroku-go"
	"github.com/spf13/cobra"
	herokuApi "github.com/tacohole/gouki/util/heroku-api"
)

var AppsCmd = &cobra.Command{
	Use:   "apps",
	Short: "get info about all apps",
	Long: `get information about all the Heroku apps you own
OPTIONS
-A, --all          include apps in all teams
-p, --personal     list apps in personal account when a default team is set
-s, --space=space  filter by space
-t, --team=team    team to use
--json             output in json format`,
	Args: cobra.MaximumNArgs(1),
	Run:  apps,
}

// var verbose bool
var allApps bool
var personalApps bool
var space string
var team string

func init() {
	AppsCmd.Flags().BoolVarP(&allApps, "all", "a", false, "include apps in all teams")
	AppsCmd.Flags().BoolVarP(&personalApps, "personal", "p", false, "list apps in personal account when a default team is set")
	AppsCmd.Flags().StringVarP(&space, "space", "s", "", "filter by space")
	AppsCmd.Flags().StringVarP(&team, "team", "t", "", "team to use")

}

func apps(cmd *cobra.Command, args []string) {
	loadDefaultVariables()
	// setup client
	client := herokuApi.MakeClient()

	// get apps
	resp, err := client.AppList(nil)
	if err != nil {
		fmt.Printf("error retrieving apps: %s", err)
	}

	appPrinter(resp, client.Username)

}

func appPrinter(apps []heroku.App, username string) {
	// print that ish
	fmt.Printf("=== %s Apps\n", username)
	for _, app := range apps {
		fmt.Printf("%s \n", app.Name)
	}
	fmt.Printf("\n")
}

func loadDefaultVariables() {
	// verbose := viper.GetBool("VERBOSE")
}
