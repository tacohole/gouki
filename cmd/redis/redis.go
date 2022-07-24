package redis

import (
	"github.com/spf13/cobra"
)

var RedisCmd = &cobra.Command{
	Use:   "redis",
	Short: "",
	Long:  ``,
	Args:  cobra.MaximumNArgs(1),
	Run:   redis,
}

var appName string

func init() {
	RedisCmd.Flags().StringVarP(&appName, "app", "a", "", "")

}

func redis(cmd *cobra.Command, args []string) {

}
