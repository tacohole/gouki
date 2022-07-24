package ci

import (
	"github.com/spf13/cobra"
)

var CiCmd = &cobra.Command{
	Use:   "ci",
	Short: "",
	Long:  ``,
	Args:  cobra.MaximumNArgs(1),
	Run:   ci,
}

// var verbose bool
// var allApps bool
// var personalApps bool
// var space string
// var team string
var appName string

func init() {
	CiCmd.Flags().StringVarP(&appName, "app", "a", "", "")

}

func ci(cmd *cobra.Command, args []string) {

}
