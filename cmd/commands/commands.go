package commands

import (
	"github.com/spf13/cobra"
)

var CommandsCmd = &cobra.Command{
	Use:   "commands",
	Short: "",
	Long:  ``,
	Args:  cobra.MaximumNArgs(1),
	Run:   commands,
}

func init() {

}

func commands(cmd *cobra.Command, args []string) {

}
