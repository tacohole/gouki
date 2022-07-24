package pipelines

import (
	"github.com/spf13/cobra"
)

var PipelinesCmd = &cobra.Command{
	Use:   "pipelines",
	Short: "",
	Long:  ``,
	Args:  cobra.MaximumNArgs(1),
	Run:   pipelines,
}

var appName string

func init() {
	PipelinesCmd.Flags().StringVarP(&appName, "app", "a", "", "")

}

func pipelines(cmd *cobra.Command, args []string) {

}
