package heroku

import (
	"os"

	"github.com/bgentry/heroku-go"
)

func MakeClient() (heroku.Client, error) {
	// headers := http.Header{}
	HEROKU_API_TOKEN := os.Getenv("HEROKU_API_TOKEN")

	// bearerToken := fmt.Sprintf("Bearer: %s", HEROKU_API_TOKEN)

	// headers.Add("Accept", "application/vnd.heroku+json; version=3")
	// headers.Add("Authorization", bearerToken)

	client := heroku.Client{
		Username: "troy.coll@gmail.com",
		Password: HEROKU_API_TOKEN}

	return client, nil
}
