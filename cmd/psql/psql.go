package psql

import (
	"github.com/spf13/cobra"
)

var PsqlCmd = &cobra.Command{
	Use:   "psql",
	Short: "",
	Long:  ``,
	Args:  cobra.MaximumNArgs(1),
	Run:   psql,
}

var appName string

func init() {
	PsqlCmd.Flags().StringVarP(&appName, "app", "a", "", "")

}

func psql(cmd *cobra.Command, args []string) {

}
