package vanilla

import (
	"fmt"
	_ "github.com/GoogleCloudPlatform/berglas/pkg/auto"
	"github.com/alexandre-normand/marcopoller"
	"net/http"
	"os"
	"time"
)

// Berglas secret environment variables
const (
	slackTokenEnv        = "SLACK_TOKEN"
	signingSecretEnv     = "SIGNING_SECRET"
	pollValidityDuration = time.Duration(24) * time.Hour
)

var mp *marcopoller.MarcoPoller

func init() {
	mpoller, err := marcopoller.NewWithOptions(marcopoller.OptionSlackVerifier(os.Getenv(signingSecretEnv)),
		marcopoller.OptionSlackClient(os.Getenv(slackTokenEnv), false),
		marcopoller.OptionDatastore(os.Getenv(marcopoller.GCPProjectIDEnv)),
		marcopoller.OptionPollVerifier(marcopoller.ExpirationPollVerifier{ValidityPeriod: pollValidityDuration}),
		marcopoller.OptionDebug(true))
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
