package auth

import (
	"log"

	"github.com/pkg/browser"
	"github.com/spf13/cobra"
)

var loginCmd = &cobra.Command{
	Short: "logs in the user",
	Long:  "logs in and writes an auth token to someplace",
	Use:   "login",
	Run:   login,
}

// -e, --expires-in=expires-in  duration of token in seconds (default 30 days)
//   -i, --interactive            login with username/password
//   --browser=browser            browser to open SSO with (example: "firefox", "safari")

func init() {
	AuthCmd.AddCommand(loginCmd)
}

// URL: https://cli-auth.heroku.com/auth/cli/browser/
// https://cli-auth.heroku.com/auth/cli/callback?code=47239e06-7781-411c-8df6-867cac68347c&state=941d093a-058c-4933-9bbe-7d72a9a52cb8

func login(cmd *cobra.Command, args []string) {
	loadDefaultVariables()
	// read the flags for browser open or interactive
	// if broswer open
	// press any key to open the browser
	// or press q to exit
	// open the browser
	const url = "https://cli-auth.heroku.com/auth/cli/browser/"
	err := browser.OpenURL(url)
	if err != nil {
		log.Fatalf("unable to open browser: %s", err)
	}

	// else grab the username
	// and the password in a non-cleartext prompt from the cli

	// take the token response and write it to .netrc

	// `Logged in as ${color.green(account.email!)}`
}
