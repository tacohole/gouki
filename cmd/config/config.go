package config

import (
	"github.com/spf13/cobra"
)

var ConfigCmd = &cobra.Command{
	Use:   "config",
	Short: "",
	Long:  ``,
	Args:  cobra.MaximumNArgs(1),
	Run:   config,
}

var appName string

func init() {
	ConfigCmd.Flags().StringVarP(&appName, "app", "a", "", "")

}

func config(cmd *cobra.Command, args []string) {

}
