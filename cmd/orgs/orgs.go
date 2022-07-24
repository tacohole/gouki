package orgs

import (
	"github.com/spf13/cobra"
)

var OrgsCmd = &cobra.Command{
	Use:   "orgs",
	Short: "",
	Long:  ``,
	Args:  cobra.MaximumNArgs(1),
	Run:   orgs,
}

var appName string

func init() {
	OrgsCmd.Flags().StringVarP(&appName, "app", "a", "", "")

}

func orgs(cmd *cobra.Command, args []string) {

}
