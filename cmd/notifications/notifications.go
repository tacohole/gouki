package notifications

import (
	"github.com/spf13/cobra"
)

var NotificationsCmd = &cobra.Command{
	Use:   "notifications",
	Short: "",
	Long:  ``,
	Args:  cobra.MaximumNArgs(1),
	Run:   notifications,
}

var appName string

func init() {
	NotificationsCmd.Flags().StringVarP(&appName, "app", "a", "", "")

}

func notifications(cmd *cobra.Command, args []string) {

}
