# 🍦 Marco Poller Vanilla (berglas-backed secrets)

[Marco Poller](https://github.com/alexandre-normand/marcopoller) is a [slack](https://slack.com) app that does one thing: manage slack polls. Yes, there are many slack polling applications at the moment. Chances are
that you don't need this. But if you're part of an organisation that doesn't allow 3rd party slack apps or if you just feel like having your own version that does its
thing a little differently, go ahead and look at [Marco Poller](https://github.com/alexandre-normand/marcopoller). This repository is a vanilla version that uses [berglas](https://github.com/GoogleCloudPlatform/berglas) to manage secrets.
This might be suitable for some deployments but, if you're in an enterprise setting with a different solution for secrets-management, you'll want to write your wrapper and use [Marco Poller](https://github.com/alexandre-normand/marcopoller).

## Backing Technology
*   [Go](https://golang.org): `>= 1.12` 
*   [Google Cloud Functions](https://cloud.google.com/functions/docs/quickstart#functions-deploy-command-go)
*   [Google Cloud Datastore](https://cloud.google.com/datastore/docs/)

## Requirements

*   A bot slack token 
    *   With the following `scopes`:
    	*   `chat:write:bot`
    	*   `chat:write:user`
    	*   `bot`
    	*   `commands`
    	*   `users.profile:read`

	*   The following `slash` commands:
    	*   `/poll`:
    	    * *Command*: `/poll`
    	    * *Request URL*: `<url of the startPoll gcloud function>`
    	    * *Short Description*: `Starts a new poll`
    	    * *Usage Hint*: `"Question?" "Option1" "Option 2"`

    *   The following `interactive` components (should be toggled to `on`):
    	*   `registerVote` action URL: This is going to show up in the `gcloud functions deploy` output for the `registerVote` function. You only have to do this when
    	     you first deploy the `registerVote` function but you'll have to enter the `URL` of the `registerVote` gcloud function in _Request URL_. 

    *	A `bot user` with the recommended info:
		*   *Display name*: `Marco Poller`
		*   *Display username*: `marcopoller`

*   A gcloud project ID with the datastore API enabled

## Integration
This ready-to-deploy vanilla version of `Marco Poller` uses [berglas](https://github.com/GoogleCloudPlatform/berglas) to manage secrets. 
Refer to the [berglas gcloud functions example and documentation](https://github.com/GoogleCloudPlatform/berglas/tree/master/examples/cloudfunctions/go) on how to set 
up berglas with gcloud functions. 

## Deploy

* Make sure you've written the `slacktoken` and slack `signingsecret` using `berglas` and that you've granted the service account access to those. If you haven't done so already, 
refer to the [berglas gcloud functions example and documentation](https://github.com/GoogleCloudPlatform/berglas/tree/master/examples/cloudfunctions/go) for how to do this.

* Deploy the vanilla/berglas version using the [gcloud cli](https://cloud.google.com/sdk/gcloud/) commands from [github.com/alexandre-normand/marcopoller/berglas](berglas):

```
gcloud functions deploy startPoll --entry-point StartPoll --runtime go111 --trigger-http --project $PROJECT_ID --service-account ${SA_EMAIL} --set-env-vars "PROJECT_ID=${PROJECT_ID},SLACK_TOKEN=berglas://${BUCKET_ID}/slacktoken,SIGNING_SECRET=berglas://${BUCKET_ID}/signingsecret"
```

```
gcloud functions deploy registerVote --entry-point RegisterVote --runtime go111 --trigger-http --project $PROJECT_ID --service-account ${SA_EMAIL} --set-env-vars "PROJECT_ID=${PROJECT_ID},SLACK_TOKEN=berglas://${BUCKET_ID}/slacktoken,SIGNING_SECRET=berglas://${BUCKET_ID}/signingsecret"
```