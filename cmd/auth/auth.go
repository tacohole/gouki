package auth

import (
	"github.com/spf13/cobra"
)

var AuthCmd = &cobra.Command{
	Use:   "auth",
	Short: "",
	Long:  "",
	Args:  cobra.MaximumNArgs(1),
}

// var verbose bool

func loadDefaultVariables() {
	// verbose = viper.GetBool("VERBOSE")
}

// https://pkg.go.dev/github.com/markbates/goth/providers/heroku
