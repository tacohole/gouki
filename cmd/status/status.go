package status

import (
	"github.com/spf13/cobra"
)

var StatusCmd = &cobra.Command{
	Use:   "status",
	Short: "",
	Long:  ``,
	Args:  cobra.MaximumNArgs(1),
	Run:   status,
}

var appName string

func init() {
	StatusCmd.Flags().StringVarP(&appName, "app", "a", "", "")

}

func status(cmd *cobra.Command, args []string) {

}
