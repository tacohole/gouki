package herokuApi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/bgentry/heroku-go"
	"github.com/tacohole/gouki/util/httpClient"
)

func MakeClient() *heroku.Client {
	HEROKU_API_TOKEN := os.Getenv("HEROKU_API_TOKEN")
	HEROKU_USERNAME := os.Getenv("HEROKU_USERNAME")

	client := heroku.Client{
		Username: HEROKU_USERNAME,
		Password: HEROKU_API_TOKEN}

	return &client
}

func MakeOAuthClient(user string, password string) (*heroku.OAuthClient, error) {
	postUri := "https://api.heroku.com/oauth/authorizations"

	headers := map[string]string{
		"Accept":        "application/vnd.heroku+json; version=3",
		"Authorization": fmt.Sprintf("Basic %s=%s", user, password),
		"Content-Type":  "application/json",
	}

	resp, err := httpClient.MakeHttpRequest("POST", postUri, headers, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	r, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var oAuth heroku.OAuthClient

	if err = json.Unmarshal(r, &oAuth); err != nil {
		return nil, err
	}

	return &oAuth, nil
}

func MakeLogSessionOpts(dyno string, num int, source string, tail bool) *heroku.LogSessionCreateOpts {
	logOpts := heroku.LogSessionCreateOpts{
		Dyno:   &dyno,
		Lines:  &num,
		Source: &source,
		Tail:   &tail,
	}

	return &logOpts
}

func MakeCollaboratorOpts(silent bool) *heroku.CollaboratorCreateOpts {
	collabOpts := heroku.CollaboratorCreateOpts{
		Silent: &silent,
	}

	return &collabOpts
}

func IsTeamApp(ownerEmail string) bool {
	return strings.Contains(ownerEmail, ".@herokumanager.com")
}

func GetTeamNameFromEmail(email string) string {
	return strings.TrimSuffix(email, ".@herokumanager.com")
}

func GetTeamInfo(teamName string) (*heroku.Organization, error) {
	team := heroku.Organization{}

	client := MakeClient()
	path := fmt.Sprintf("/teams/%s/features", teamName)

	if err := client.APIReq(team, "GET", path, nil); err != nil {
		return nil, err
	}

	return &team, nil
}

func AddTeamCollaborator(app string, email string, permissions []string, silent bool) (*heroku.OrganizationAppCollaborator, error) {
	collab := heroku.OrganizationAppCollaborator{}

	client := MakeClient()
	path := fmt.Sprintf("/teams/apps/%s/collaborators", app)

	body := TeamCollaborator{
		permissions: permissions,
		email:       email,
		silent:      silent,
	}

	if err := client.APIReq(collab, "POST", path, body); err != nil {
		return nil, err
	}

	return &collab, nil
}

func UpdateTeamCollaborator(app string, email string, permissions []string) (*heroku.OrganizationAppCollaborator, error) {
	collab := heroku.OrganizationAppCollaborator{}

	client := MakeClient()
	path := fmt.Sprintf("/teams/apps/%s/collaborators/%s", app, email)

	if err := client.APIReq(collab, "PATCH", path, permissions); err != nil {
		return nil, err
	}

	return &collab, nil
}
