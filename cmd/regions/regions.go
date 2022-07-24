package regions

import (
	"github.com/spf13/cobra"
)

var RegionsCmd = &cobra.Command{
	Use:   "regions",
	Short: "",
	Long:  ``,
	Args:  cobra.MaximumNArgs(1),
	Run:   regions,
}

var appName string

func init() {
	RegionsCmd.Flags().StringVarP(&appName, "app", "a", "", "")

}

func regions(cmd *cobra.Command, args []string) {

}
