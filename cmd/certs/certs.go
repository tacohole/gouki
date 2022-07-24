package certs

import (
	"github.com/spf13/cobra"
)

var CertsCmd = &cobra.Command{
	Use:   "certs",
	Short: "",
	Long:  ``,
	Args:  cobra.MaximumNArgs(1),
	Run:   certs,
}

// var verbose bool
// var allApps bool
// var personalApps bool
// var space string
// var team string
var appName string

func init() {
	CertsCmd.Flags().StringVarP(&appName, "app", "a", "", "")

}

func certs(cmd *cobra.Command, args []string) {

}
