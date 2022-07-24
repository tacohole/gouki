package reviewapps

import (
	"github.com/spf13/cobra"
)

var ReviewAppsCmd = &cobra.Command{
	Use:   "reviewapps",
	Short: "",
	Long:  ``,
	Args:  cobra.MaximumNArgs(1),
	Run:   reviewapps,
}

var appName string

func init() {
	ReviewAppsCmd.Flags().StringVarP(&appName, "app", "a", "", "")

}

func reviewapps(cmd *cobra.Command, args []string) {

}
