package access

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	herokuApi "github.com/tacohole/gouki/util/heroku-api"
)

var AccessCmd = &cobra.Command{
	Short: "list who has access to an app",
	Long: `USAGE
  $ gouki access -a APPNAME`,
	Use:  "access",
	Args: cobra.MaximumNArgs(1),
	Run:  accessInfo,
}

var user string
var permissions []string

type AppAccess struct {
	Email string
	Role  string
}

func init() {
	AccessCmd.MarkFlagRequired("app")

}

func accessInfo(cmd *cobra.Command, args []string) {
	app := viper.GetString("app")
	if app == "" {
		log.Fatalf("app name is required for this command")
	}

	client := herokuApi.MakeClient()

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
	collaborators, err := client.CollaboratorList(app, nil)
	if err != nil {
		log.Printf("%v", err)
	}

	if herokuApi.IsTeamApp(ownerEmail) {
		// first get the team name from the email
		teamName := herokuApi.GetTeamNameFromEmail(ownerEmail)

		members, err := client.OrganizationMemberList(teamName, nil)
		if err != nil {
			log.Printf("%v", err)
		}

		for _, m := range members {
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

	accessPrinter(users)

}

func accessPrinter(users []AppAccess) {
	// TODO: some kind of a sort here to display things with better spacing
	for _, u := range users {
		fmt.Printf("%s  %s \n", u.Email, u.Role)
	}
}
