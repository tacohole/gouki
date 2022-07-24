package features

import (
	"github.com/spf13/cobra"
)

var FeaturesCmd = &cobra.Command{
	Use:   "features",
	Short: "",
	Long:  ``,
	Args:  cobra.MaximumNArgs(1),
	Run:   features,
}

var appName string

func init() {
	FeaturesCmd.Flags().StringVarP(&appName, "app", "a", "", "")

}

func features(cmd *cobra.Command, args []string) {

}
