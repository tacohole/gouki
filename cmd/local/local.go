package local

import (
	"github.com/spf13/cobra"
)

var LocalCmd = &cobra.Command{
	Use:   "local",
	Short: "",
	Long:  ``,
	Args:  cobra.MaximumNArgs(1),
	Run:   local,
}

var appName string

func init() {
	LocalCmd.Flags().StringVarP(&appName, "app", "a", "", "")

}

func local(cmd *cobra.Command, args []string) {

}
