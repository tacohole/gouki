package keys

import (
	"github.com/spf13/cobra"
)

var KeysCmd = &cobra.Command{
	Use:   "keys",
	Short: "",
	Long:  ``,
	Args:  cobra.MaximumNArgs(1),
	Run:   keys,
}

var appName string

func init() {
	KeysCmd.Flags().StringVarP(&appName, "app", "a", "", "")

}

func keys(cmd *cobra.Command, args []string) {

}
