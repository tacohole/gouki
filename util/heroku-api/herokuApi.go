package herokuApi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

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
