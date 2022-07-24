package pg

import (
	"github.com/spf13/cobra"
)

var PgCmd = &cobra.Command{
	Use:   "pg",
	Short: "",
	Long:  ``,
	Args:  cobra.MaximumNArgs(1),
	Run:   pg,
}

var appName string

func init() {
	PgCmd.Flags().StringVarP(&appName, "app", "a", "", "")

}

func pg(cmd *cobra.Command, args []string) {

}
