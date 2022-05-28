package heroku

import (
	"github.com/bgentry/heroku-go"
)

func makeClient() (heroku.Client, error) {
	client := heroku.Client{Username: "email@me.com", Password: "my-api-key"}

	return client, nil
}
