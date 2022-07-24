package builds

import (
	"github.com/spf13/cobra"
)

var BuildsCmd = &cobra.Command{
	Use:   "builds",
	Short: "",
	Long:  ``,
	Args:  cobra.MaximumNArgs(1),
	Run:   builds,
}

// var verbose bool
// var allApps bool
// var personalApps bool
// var space string
// var team string
var appName string

func init() {
	BuildsCmd.Flags().StringVarP(&appName, "app", "a", "", "")

}

func builds(cmd *cobra.Command, args []string) {

}
