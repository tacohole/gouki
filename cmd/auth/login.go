package auth

import (
	"log"

	"github.com/pkg/browser"
	"github.com/spf13/cobra"
)

var loginCmd = &cobra.Command{
	Short: "logs in the user",
	Long:  "logs in and writes an auth token tom someplace",
	Use:   "login",
	Run:   login,
}

func init() {
	AuthCmd.AddCommand(loginCmd)
}

func login(cmd *cobra.Command, args []string) {
	// read the flags for browser open or interactive
	// if broswer open
	// press any key to open the browser
	// or press q to exit
	// open the browser
	const url = "https://id.heroku.com"
	err := browser.OpenURL(url)
	if err != nil {
		log.Fatalf("unable to open browser: %s", err)
	}

	// else grab the username
	// and the password in a non-cleartext prompt from the cli

	// take the token response and write it to .netrc
}
