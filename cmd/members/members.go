package members

import (
	"github.com/spf13/cobra"
)

var MembersCmd = &cobra.Command{
	Use:   "members",
	Short: "",
	Long:  ``,
	Args:  cobra.MaximumNArgs(1),
	Run:   members,
}

var appName string

func init() {
	MembersCmd.Flags().StringVarP(&appName, "app", "a", "", "")

}

func members(cmd *cobra.Command, args []string) {

}
