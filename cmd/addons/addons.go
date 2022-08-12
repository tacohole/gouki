package addons

import (
	"fmt"
	"log"

	"github.com/bgentry/heroku-go"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	herokuApi "github.com/tacohole/gouki/util/heroku-api"
)

var AddonsCmd = &cobra.Command{
	Short: "shows addons",
	Long:  "",
	Use:   "addons",
	Run:   addons,
}

func init() {

}

func addons(cmd *cobra.Command, args []string) {
	appName := viper.GetString("app")

	// check our current directory
	// if the current directory contains a heroku remote
	// use that app name as the appName

	if appName != "" {
		addonsList, err := addonsByApp(appName)
		if err != nil {
			log.Printf("%v", err)
		}
		addonsPrinter(addonsList, appName)
	} else {
		addonsList, err := herokuApi.GetAllAddons()
		if err != nil {
			log.Printf("%v", err)
		}
		addonsPrinter(*addonsList, appName)
	}
}

func addonsByApp(appName string) ([]heroku.Addon, error) {
	client := herokuApi.MakeClient()

	addons, err := client.AddonList(appName, nil)
	if err != nil {
		return nil, err
	}

	return addons, nil
}

func addonsPrinter(addons []heroku.Addon, appName string) {
	fmt.Printf("Owning App    Add-on    Plan    Price      State\n")
	for _, addon := range addons {
		fmt.Printf("%s    %s    %s    %s    %s\n", appName, addon.Name, addon.Plan.Name, "0", "created")
	}
}
