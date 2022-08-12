package access

import (
	"testing"

	"github.com/bgentry/heroku-go"
)

var indApp = heroku.App{
	Name: "personal-app",
	Owner: struct {
		Email string "json:\"email\""
		Id    string "json:\"id\""
	}{
		Email: "personal@appowner.com",
		Id:    "id-for-personal-app",
	},
}

var teamApp = heroku.App{
	Name: "team-app",
	Owner: struct {
		Email string "json:\"email\""
		Id    string "json:\"id\""
	}{
		Email: "teamname@herokumanager.com",
		Id:    "id-for-team-app",
	},
}

var collaborator1 = heroku.Collaborator{
	Id: "id-for-personal-app-collaborator",
	User: struct {
		Email string "json:\"email\""
		Id    string "json:\"id\""
	}{
		Email: "collaborator@appcollaborator.com",
		Id:    "id-for-collab-user-1",
	},
}

var collaborator2 = heroku.Collaborator{
	Id: "id-for-team-app-collaborator",
	User: struct {
		Email string "json:\"email\""
		Id    string "json:\"id\""
	}{
		Email: "collaborator2@appcollaborator.com",
		Id:    "id-for-collab-user-2",
	},
}

// var orgMember1 = heroku.OrganizationMember{

// }

// var orgMember2 = heroku.OrganizationMember{}

// print access
func TestAccessInfo(t *testing.T) {
	// return user names on app
}

// add user
func TestAccessAdd(t *testing.T) {
	// run accessAdd with app + new collaborator
	// expect accessInfo to return new collaborator
	// run accessAdd with app + new org member + permissions
	// expect accessInfo to return new org member + permissions
}

// remove user
func TestAccessRemove(t *testing.T) {}

// update user
func TestAccessUpdate(t *testing.T) {}
