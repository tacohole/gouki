package heroku

import (
	"os"

	"github.com/bgentry/heroku-go"
)

var (
	HEROKU_USERNAME  = os.Getenv("HEROKU_USERNAME")
	HEROKU_API_TOKEN = os.Getenv("HEROKU_API_TOKEN")
)

func MakeClient() (heroku.Client, error) {
	client := heroku.Client{Username: HEROKU_USERNAME, Password: HEROKU_API_TOKEN}

	return client, nil
}
