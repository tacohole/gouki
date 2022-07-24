package trustedips

import (
	"github.com/spf13/cobra"
)

var TrustedIpsCmd = &cobra.Command{
	Use:   "trusted-ips",
	Short: "",
	Long:  ``,
	Args:  cobra.MaximumNArgs(1),
	Run:   trustedIps,
}

var appName string

func init() {
	TrustedIpsCmd.Flags().StringVarP(&appName, "app", "a", "", "")

}

func trustedIps(cmd *cobra.Command, args []string) {

}
