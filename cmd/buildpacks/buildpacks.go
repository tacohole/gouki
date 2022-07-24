package buildpacks

import (
	"github.com/spf13/cobra"
)

var BuildpacksCmd = &cobra.Command{
	Use:   "buildpacks",
	Short: "",
	Long:  ``,
	Args:  cobra.MaximumNArgs(1),
	Run:   buildpacks,
}

// var verbose bool
// var allApps bool
// var personalApps bool
// var space string
// var team string
var appName string

func init() {
	BuildpacksCmd.Flags().StringVarP(&appName, "app", "a", "", "")

}

func buildpacks(cmd *cobra.Command, args []string) {

}
