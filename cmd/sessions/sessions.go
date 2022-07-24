package sessions

import (
	"github.com/spf13/cobra"
)

var SessionsCmd = &cobra.Command{
	Use:   "sessions",
	Short: "",
	Long:  ``,
	Args:  cobra.MaximumNArgs(1),
	Run:   sessions,
}

var appName string

func init() {
	SessionsCmd.Flags().StringVarP(&appName, "app", "a", "", "")

}

func sessions(cmd *cobra.Command, args []string) {

}
