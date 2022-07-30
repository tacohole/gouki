package access

import (
	"log"
	"strings"

	"github.com/spf13/cobra"
	herokuApi "github.com/tacohole/gouki/util/heroku-api"
)

var AccessCmd = &cobra.Command{
	Short: "list who has access to an app",
	Long:  ``,
	Use:   "access",
	Args:  cobra.MinimumNArgs(1),
	Run:   access,
}

type AppAccess struct {
	Email string
	Role  string
}

func init() {

}

func access(cmd *cobra.Command, args []string) {
	// client
	client := herokuApi.MakeClient()

	// app get info
	AppInfo, err := client.AppInfo(appName)
	if err != nil {
		log.Printf("%v", err)
	}

	// return value
	users := []AppAccess{}

	// get app collaborators
	// collaborator.User.Email
	collaborators, err := client.CollaboratorList(appName, nil)
	if err != nil {
		log.Printf("%v", err)
	}

	ownerEmail := AppInfo.Owner.Email

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

	// pull the admins out of the collaborators array
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
		log.Printf("%s			%s \n", u.Email, u.Role)
	}
}

func getTeamNameFromEmail(email string) string {
	return strings.TrimSuffix(email, ".@herokumanager.com")
}

func isTeamApp(ownerEmail string) bool {
	// use a regexp to check if the app is owned by a Heroku team
	return strings.Contains(ownerEmail, ".@herokumanager.com")
}
