package git

import (
	"github.com/spf13/cobra"
)

var GitCmd = &cobra.Command{
	Use:   "git",
	Short: "",
	Long:  ``,
	Args:  cobra.MaximumNArgs(1),
	Run:   git,
}

var appName string

func init() {
	GitCmd.Flags().StringVarP(&appName, "app", "a", "", "")

}

func git(cmd *cobra.Command, args []string) {

}
