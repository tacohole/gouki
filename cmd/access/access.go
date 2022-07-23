package access

import "github.com/spf13/cobra"

var AccessCmd = &cobra.Command{
	Short: "log in with your Heroku credentials",
	Long:  ``,
	Use:   "access",
	Run:   access,
}

var appName string
var dyno string
var num int
var remote string
var source string
var tail bool

func init() {
	AccessCmd.Flags().StringVarP(&appName, "app", "a", "", "")
	AccessCmd.Flags().StringVarP(&dyno, "dyno", "d", "", "")
	AccessCmd.Flags().IntVarP(&num, "num", "n", 0, "")
	AccessCmd.Flags().StringVarP(&remote, "remote", "r", "", "")
	AccessCmd.Flags().StringVarP(&source, "source", "s", "", "")
	AccessCmd.Flags().BoolVarP(&tail, "tail", "t", false, "")

}

func access(cmd *cobra.Command, args []string) {}
