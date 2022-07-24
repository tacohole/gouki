package run

import (
	"github.com/spf13/cobra"
)

var RunCmd = &cobra.Command{
	Use:   "run",
	Short: "",
	Long:  ``,
	Args:  cobra.MaximumNArgs(1),
	Run:   run,
}

var appName string

func init() {
	RunCmd.Flags().StringVarP(&appName, "app", "a", "", "")

}

func run(cmd *cobra.Command, args []string) {

}
