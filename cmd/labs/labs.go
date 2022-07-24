package labs

import (
	"github.com/spf13/cobra"
)

var LabsCmd = &cobra.Command{
	Use:   "labs",
	Short: "",
	Long:  ``,
	Args:  cobra.MaximumNArgs(1),
	Run:   labs,
}

var appName string

func init() {
	LabsCmd.Flags().StringVarP(&appName, "app", "a", "", "")

}

func labs(cmd *cobra.Command, args []string) {

}
