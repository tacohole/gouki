package drains

import (
	"github.com/spf13/cobra"
)

var DrainsCmd = &cobra.Command{
	Use:   "drains",
	Short: "",
	Long:  ``,
	Args:  cobra.MaximumNArgs(1),
	Run:   drains,
}

var appName string

func init() {
	DrainsCmd.Flags().StringVarP(&appName, "app", "a", "", "")

}

func drains(cmd *cobra.Command, args []string) {

}
