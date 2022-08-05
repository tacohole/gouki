package addons

import "github.com/spf13/cobra"

var AddonsCmd = &cobra.Command{
	Short: "shows addons",
	Long:  "",
	Use:   "addons",
	Run:   addons,
}

func init() {

}

func addons(cmd *cobra.Command, args []string) {}
