package vanilla

import (
	"fmt"
	_ "github.com/GoogleCloudPlatform/berglas/pkg/auto"
	"github.com/alexandre-normand/marcopoller"
	"net/http"
	"os"
)

// Berglas secret environment variables
const (
	slackTokenEnv    = "SLACK_TOKEN"
	signingSecretEnv = "SIGNING_SECRET"
)

var mp *marcopoller.MarcoPoller

func init() {
	mpoller, err := marcopoller.New(os.Getenv(slackTokenEnv), os.Getenv(signingSecretEnv), os.Getenv(marcopoller.GCPProjectIDEnv))
	if err != nil {
		panic(fmt.Sprintf("Failed to initialize Marco Poller: %s", err.Error()))
	}

	mp = mpoller
}

// StartPoll handles a slash command request to start a new poll using berglas backed secrets
func StartPoll(w http.ResponseWriter, r *http.Request) {
	mp.StartPoll(w, r)
}

// RegisterVote handles a slash command request to register a poll vote using berglas backed secrets
func RegisterVote(w http.ResponseWriter, r *http.Request) {
	mp.RegisterVote(w, r)
}
