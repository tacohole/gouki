package heroku

import (
	"os"

	"github.com/bgentry/heroku-go"
)

func MakeClient() (heroku.Client, error) {
	HEROKU_API_TOKEN := os.Getenv("HEROKU_API_TOKEN")
	HEROKU_USERNAME := os.Getenv("HEROKU_USERNAME")

	client := heroku.Client{
		Username: HEROKU_USERNAME,
		Password: HEROKU_API_TOKEN}

	return client, nil
}
