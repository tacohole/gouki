package access

import (
	"fmt"
	"log"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	herokuApi "github.com/tacohole/gouki/util/heroku-api"
)

var AccessCmd = &cobra.Command{
	Short: "list who has access to an app",
	Long: `USAGE
  $ gouki access

OPTIONS
  -a, --app=app        (required) app to run command against
  -r, --remote=remote  git remote of app to use
  --json               output in json format`,
	Use:  "access",
	Args: cobra.MaximumNArgs(1),
	Run:  access,
}

type AppAccess struct {
	Email string
	Role  string
}

func init() {
}

func access(cmd *cobra.Command, args []string) {
	app := viper.GetString("app")

	// client
	client := herokuApi.MakeClient()

	// app get info
	AppInfo, err := client.AppInfo(app)
	if err != nil {
		log.Printf("%v", err)
	}

	// return value
	users := []AppAccess{}

	ownerEmail := AppInfo.Owner.Email
	appOwner := AppAccess{
		Email: ownerEmail,
		Role:  "owner",
	}
	users = append(users, appOwner)

	// get app collaborators
	// collaborator.User.Email
	collaborators, err := client.CollaboratorList(app, nil)
	if err != nil {
		log.Printf("%v", err)
	}

	// check if the app is owned by a team
	if isTeamApp(ownerEmail) {
		// first get the team name from the email
		teamName := getTeamNameFromEmail(ownerEmail)

		// then get the team members
		// member.Email, member.Role
		members, err := client.OrganizationMemberList(teamName, nil)
		if err != nil {
			log.Printf("%v", err)
		}

		// since all team admins have access to all apps, pick out the admins
		for _, m := range members {
			// if the role is admin, add them to the []AppAccess
			if m.Role == "admin" {
				user := AppAccess{
					Email: m.Email,
					Role:  m.Role,
				}
				users = append(users, user)
			}
		}
	}

	// skip the admins in the collaborators array
	for _, u := range users {
		for _, c := range collaborators {
			if c.User.Email == u.Email {
				continue
			}
			user := AppAccess{
				Email: c.User.Email,
				Role:  "collaborator",
			}
			users = append(users, user)
		}
	}

	// print it
	accessPrinter(users)

}

func accessPrinter(users []AppAccess) {
	for _, u := range users {
		fmt.Printf("%s  %s \n", u.Email, u.Role)
	}
}

func getTeamNameFromEmail(email string) string {
	return strings.TrimSuffix(email, ".@herokumanager.com")
}

func isTeamApp(ownerEmail string) bool {
	// use a regexp to check if the app is owned by a Heroku team
	return strings.Contains(ownerEmail, ".@herokumanager.com")
}
