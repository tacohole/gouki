/*
Copyright © 2022 TROY COLL <troy.coll@gmail.com>

*/
package cmd

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"github.com/tacohole/gouki/cmd/access"
	"github.com/tacohole/gouki/cmd/apps"
	"github.com/tacohole/gouki/cmd/auth"
	"github.com/tacohole/gouki/cmd/buildpacks"
	"github.com/tacohole/gouki/cmd/builds"
	"github.com/tacohole/gouki/cmd/certs"
	"github.com/tacohole/gouki/cmd/ci"
	"github.com/tacohole/gouki/cmd/commands"
	"github.com/tacohole/gouki/cmd/config"
	"github.com/tacohole/gouki/cmd/container"
	"github.com/tacohole/gouki/cmd/domains"
	"github.com/tacohole/gouki/cmd/drains"
	"github.com/tacohole/gouki/cmd/dyno"
	"github.com/tacohole/gouki/cmd/features"
	"github.com/tacohole/gouki/cmd/git"
	"github.com/tacohole/gouki/cmd/keys"
	"github.com/tacohole/gouki/cmd/labs"
	"github.com/tacohole/gouki/cmd/local"
	"github.com/tacohole/gouki/cmd/logs"
	"github.com/tacohole/gouki/cmd/maintenance"
	"github.com/tacohole/gouki/cmd/members"
	"github.com/tacohole/gouki/cmd/notifications"
	"github.com/tacohole/gouki/cmd/orgs"
	"github.com/tacohole/gouki/cmd/pg"
	"github.com/tacohole/gouki/cmd/pipelines"
	"github.com/tacohole/gouki/cmd/ps"
	"github.com/tacohole/gouki/cmd/psql"
	"github.com/tacohole/gouki/cmd/redis"
	"github.com/tacohole/gouki/cmd/regions"
	"github.com/tacohole/gouki/cmd/releases"
	"github.com/tacohole/gouki/cmd/reviewapps"
	"github.com/tacohole/gouki/cmd/run"
	"github.com/tacohole/gouki/cmd/sessions"
	"github.com/tacohole/gouki/cmd/spaces"
	"github.com/tacohole/gouki/cmd/status"
	trustedips "github.com/tacohole/gouki/cmd/trusted-ips"
	"github.com/tacohole/gouki/cmd/webhooks"
)

var cfgFile string
var Verbose bool
var Json bool
var Extended bool
var AppName string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gouki",
	Short: "The Heroku CLI but in Go",
	Long:  `I said what I said`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.gouki.yaml)")
	rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")
	rootCmd.PersistentFlags().BoolVarP(&Json, "json", "j", false, "--json provides output in JSON format")
	rootCmd.PersistentFlags().BoolVarP(&Extended, "extended", "x", false, "-x provides extended information")

	rootCmd.PersistentFlags().StringVarP(&AppName, "app", "a", "", "the name of the Heroku app")

	rootCmd.AddCommand(access.AccessCmd)
	// rootCmd.AddCommand(addons.AddonsCmd)
	rootCmd.AddCommand(apps.AppsCmd)
	rootCmd.AddCommand(auth.AuthCmd)
	rootCmd.AddCommand(buildpacks.BuildpacksCmd)
	rootCmd.AddCommand(builds.BuildsCmd)
	rootCmd.AddCommand(certs.CertsCmd)
	rootCmd.AddCommand(ci.CiCmd)
	rootCmd.AddCommand(commands.CommandsCmd)
	rootCmd.AddCommand(config.ConfigCmd)
	rootCmd.AddCommand(container.ContainerCmd)
	rootCmd.AddCommand(domains.DomansCmd)
	rootCmd.AddCommand(drains.DrainsCmd)
	rootCmd.AddCommand(dyno.DynoCmd)
	rootCmd.AddCommand(features.FeaturesCmd)
	rootCmd.AddCommand(git.GitCmd)
	rootCmd.AddCommand(keys.KeysCmd)
	rootCmd.AddCommand(labs.LabsCmd)
	rootCmd.AddCommand(local.LocalCmd)
	rootCmd.AddCommand(logs.LogsCmd)
	rootCmd.AddCommand(maintenance.MaintenanceCmd)
	rootCmd.AddCommand(members.MembersCmd)
	rootCmd.AddCommand(notifications.NotificationsCmd)
	rootCmd.AddCommand(orgs.OrgsCmd)
	rootCmd.AddCommand(pg.PgCmd)
	rootCmd.AddCommand(pipelines.PipelinesCmd)
	rootCmd.AddCommand(ps.PsCmd)
	rootCmd.AddCommand(psql.PsqlCmd)
	rootCmd.AddCommand(redis.RedisCmd)
	rootCmd.AddCommand(regions.RegionsCmd)
	rootCmd.AddCommand(releases.ReleasesCmd)
	rootCmd.AddCommand(reviewapps.ReviewAppsCmd)
	rootCmd.AddCommand(run.RunCmd)
	rootCmd.AddCommand(sessions.SessionsCmd)
	rootCmd.AddCommand(spaces.SpacesCmd)
	rootCmd.AddCommand(status.StatusCmd)
	rootCmd.AddCommand(trustedips.TrustedIpsCmd)
	rootCmd.AddCommand(webhooks.WebhooksCmd)

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	godotenv.Load(".env")

	// if cfgFile != "" {
	// 	// Use config file from the flag.
	// 	viper.SetConfigFile(cfgFile)
	// } else {
	// 	// Find home directory.
	// 	home, err := os.UserHomeDir()
	// 	cobra.CheckErr(err)

	// 	// Search config in home directory with name ".gouki" (without extension).
	// 	viper.AddConfigPath(home)
	// 	viper.SetConfigType("yaml")
	// 	viper.SetConfigName(".gouki")
	// }

	// viper.AutomaticEnv() // read in environment variables that match

	// // If a config file is found, read it in.
	// if err := viper.ReadInConfig(); err == nil {
	// 	fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	// }
}
