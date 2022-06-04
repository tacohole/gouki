package heroku

import (
	"github.com/bgentry/heroku-go"
)

func MakeClient() (heroku.Client, error) {
	client := heroku.Client{Username: "email@me.com", Password: "my-api-key"}

	return client, nil
}
