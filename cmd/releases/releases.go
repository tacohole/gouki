package releases

import (
	"github.com/spf13/cobra"
)

var ReleasesCmd = &cobra.Command{
	Use:   "releases",
	Short: "",
	Long:  ``,
	Args:  cobra.MaximumNArgs(1),
	Run:   releases,
}

var appName string

func init() {
	ReleasesCmd.Flags().StringVarP(&appName, "app", "a", "", "")

}

func releases(cmd *cobra.Command, args []string) {

}
