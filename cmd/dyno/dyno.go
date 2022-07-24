package dyno

import (
	"github.com/spf13/cobra"
)

var DynoCmd = &cobra.Command{
	Use:   "dyno",
	Short: "",
	Long:  ``,
	Args:  cobra.MaximumNArgs(1),
	Run:   dyno,
}

var appName string

func init() {
	DynoCmd.Flags().StringVarP(&appName, "app", "a", "", "")

}

func dyno(cmd *cobra.Command, args []string) {

}
