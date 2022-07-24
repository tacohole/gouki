package ps

import (
	"github.com/spf13/cobra"
)

var PsCmd = &cobra.Command{
	Use:   "ps",
	Short: "",
	Long:  ``,
	Args:  cobra.MaximumNArgs(1),
	Run:   ps,
}

var appName string

func init() {
	PsCmd.Flags().StringVarP(&appName, "app", "a", "", "")

}

func ps(cmd *cobra.Command, args []string) {

}
