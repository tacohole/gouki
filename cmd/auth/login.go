package auth

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"syscall"

	"github.com/bgentry/heroku-go"
	"github.com/spf13/cobra"
	"golang.org/x/term"

	herokuApi "github.com/tacohole/gouki/util/heroku-api"
)

var loginCmd = &cobra.Command{
	Short: "log in with your Heroku credentials",
	Long: `logs in and writes an auth token to someplace
	OPTIONS
-e, --expires-in=expires-in  duration of token in seconds (default 30 days)
-i, --interactive            login with username/password, not available for MFA accounts
--browser=browser            browser to open SSO with (example: "firefox", "safari")`,
	Use: "login",
	Run: login,
}

var expiresIn string
var interactive bool
var browserVar string

const Url = "https://cli-auth.heroku.com/auth/cli/browser/"

func init() {
	AuthCmd.Flags().StringVarP(&expiresIn, "expiresIn", "e", "", "")
	AuthCmd.Flags().BoolVar(&interactive, "interactive", false, "")
	AuthCmd.Flags().StringVarP(&browserVar, "browser", "", "", "")

	AuthCmd.AddCommand(loginCmd)
}

// URL: https://cli-auth.heroku.com/auth/cli/browser/
// https://cli-auth.heroku.com/auth/cli/callback?code=47239e06-7781-411c-8df6-867cac68347c&state=941d093a-058c-4933-9bbe-7d72a9a52cb8

func login(cmd *cobra.Command, args []string) {

	oAuth, err := getToken(interactive)
	if err != nil {
		log.Printf("error grabbing token: %v", err)
	}

	token := oAuth.Secret

	// take the token response and write it to .netrc
	// open .netrc

	netRc, err := os.Create("~/.netrc")
	if err != nil {
		log.Printf("can't open .netrc: %v", err)
	}
	defer netRc.Close()
	// write token to .netrc
	if _, err = io.WriteString(netRc, token); err != nil {
		log.Printf("can't write to netrc: %v", err)
	}
	netRc.Sync()

	userName, err := getUserName()
	if err != nil {
		log.Printf("can't get username: %v", err)
	}

	log.Printf("Logged in as %s!", *userName)
}

func getUserName() (*string, error) {
	// call to API for account info
	client := herokuApi.MakeClient()

	account, err := client.AccountInfo()
	if err != nil {
		log.Printf("error: %v", err)
	}

	return &account.Email, nil
}

func getToken(i bool) (*heroku.OAuthClient, error) {
	// read the flags for browser open or interactive
	if i {
		token, err := interactiveLogin()
		if err != nil {
			log.Printf("error: %v", err)
		}
		return token, nil

	} else {
		// token, err := browserLogin(browserVar)
		// if err != nil {
		// 	log.Printf("error: %v", err)
		// }
		return nil, nil
	}
}

// func browserLogin(b string) (*string, error) {
// 	// press any key to open the browser
// 	// or press q to exit
// 	// open the browser

// 	if err := browser.OpenURL(Url); err != nil {
// 		return nil, fmt.Errorf("unable to open browser: %v", err)
// 	}

// 	return nil, nil
// }

func interactiveLogin() (*heroku.OAuthClient, error) {
	// else grab the username
	// and the password in a non-cleartext prompt from the cli
	reader := bufio.NewReader(os.Stdin)

	log.Printf("username: ")
	rawUser, err := reader.ReadString('\n')
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	user := strings.TrimSpace(rawUser)

	log.Printf("password: ")
	bytePw, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	password := string(bytePw)

	// make our own OAuth client with the username and password
	oAuth, err := herokuApi.MakeOAuthClient(user, password)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	return oAuth, nil
}

// API call to actually do the token thing for interactive
