package spaces

import (
	"github.com/spf13/cobra"
)

var SpacesCmd = &cobra.Command{
	Use:   "spaces",
	Short: "",
	Long:  ``,
	Args:  cobra.MaximumNArgs(1),
	Run:   spaces,
}

var appName string

func init() {
	SpacesCmd.Flags().StringVarP(&appName, "app", "a", "", "")

}

func spaces(cmd *cobra.Command, args []string) {

}
