package logs

import (
	"io"
	"log"
	"os"
	"time"

	"github.com/spf13/cobra"
	herokuApi "github.com/tacohole/gouki/util/heroku-api"
	"github.com/tacohole/gouki/util/httpClient"
)

var LogsCmd = &cobra.Command{
	Short: "",
	Long:  ``,
	Use:   "logs",
	Run:   logs,
}

var appName string
var dyno string
var num int
var remote string
var source string
var tail bool

func init() {
	LogsCmd.Flags().StringVarP(&appName, "app", "a", "", "")
	LogsCmd.Flags().StringVarP(&dyno, "dyno", "d", "", "")
	LogsCmd.Flags().IntVarP(&num, "num", "n", 50, "")
	LogsCmd.Flags().StringVarP(&remote, "remote", "r", "", "")
	LogsCmd.Flags().StringVarP(&source, "source", "s", "", "")
	LogsCmd.Flags().BoolVarP(&tail, "tail", "t", false, "")

}

func logs(cmd *cobra.Command, args []string) {
	client := herokuApi.MakeClient()

	// logSessionCreateOpts
	opts := herokuApi.MakeLogSessionOpts(dyno, num, source, tail)

	session, err := client.LogSessionCreate(appName, opts)
	if err != nil {
		log.Printf("%v", err)
	}

	url := session.LogplexURL
	resp, err := httpClient.MakeHttpRequest("GET", url, nil, nil)
	if err != nil {
		log.Printf("%v", err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)

	if tail {
		for {
			time.Sleep(5 * time.Second)
			resp, err := httpClient.MakeHttpRequest("GET", url, nil, nil)
			if err != nil {
				log.Printf("%v", err)
			}

			io.Copy(os.Stdout, resp.Body)
		}
	}
}
