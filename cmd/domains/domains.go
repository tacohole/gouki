package domains

import (
	"github.com/spf13/cobra"
)

var DomansCmd = &cobra.Command{
	Use:   "domains",
	Short: "",
	Long:  ``,
	Args:  cobra.MaximumNArgs(1),
	Run:   domains,
}

var appName string

func init() {
	DomansCmd.Flags().StringVarP(&appName, "app", "a", "", "")

}

func domains(cmd *cobra.Command, args []string) {

}
