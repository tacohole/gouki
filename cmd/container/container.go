package container

import (
	"github.com/spf13/cobra"
)

var ContainerCmd = &cobra.Command{
	Use:   "container",
	Short: "",
	Long:  ``,
	Args:  cobra.MaximumNArgs(1),
	Run:   container,
}

var appName string

func init() {
	ContainerCmd.Flags().StringVarP(&appName, "app", "a", "", "")

}

func container(cmd *cobra.Command, args []string) {

}
