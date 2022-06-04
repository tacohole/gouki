package apps

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tacohole/gouki/util/heroku"
)

var AppsCmd = &cobra.Command{
	Use:   "apps",
	Short: "",
	Long:  "",
	Args:  cobra.MaximumNArgs(1),
	Run:   apps,
}

// var verbose bool
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
		fmt.Printf("%s", client.AdditionalHeaders.Get("Authorization"))
	}

	// print that ish
	for _, app := range resp {
		fmt.Printf("%s", app.Name)
	}

}

func loadDefaultVariables() {
	// verbose = viper.GetBool("VERBOSE")
}
