package maintenance

import (
	"github.com/spf13/cobra"
)

var MaintenanceCmd = &cobra.Command{
	Use:   "maintenance",
	Short: "",
	Long:  ``,
	Args:  cobra.MaximumNArgs(1),
	Run:   maintenance,
}

var appName string

func init() {
	MaintenanceCmd.Flags().StringVarP(&appName, "app", "a", "", "")

}

func maintenance(cmd *cobra.Command, args []string) {

}
