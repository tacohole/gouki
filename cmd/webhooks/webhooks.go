package webhooks

import (
	"github.com/spf13/cobra"
)

var WebhooksCmd = &cobra.Command{
	Use:   "webhooks",
	Short: "",
	Long:  ``,
	Args:  cobra.MaximumNArgs(1),
	Run:   webhooks,
}

var appName string

func init() {
	WebhooksCmd.Flags().StringVarP(&appName, "app", "a", "", "")

}

func webhooks(cmd *cobra.Command, args []string) {

}
