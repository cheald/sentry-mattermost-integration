# sentry-mattermost-sidecar

This tools is a sidecar to use Sentry internal integrations with Mattermost.

Right now it only handles the "new issue triggered" event.

## How to use

First you must create a [Mattermost incoming webhook](https://docs.mattermost.com/developer/webhooks-incoming.html) integration.
![mattermost-incoming-webhook-integration-setup](docs/assets/mattermost-incoming-webhook-integration-setup.png)

Next you must deploy the [docker image](https://hub.docker.com/r/itsalex/sentry-mattermost-sidecar) (don't forget to fill `SMS_MATTERMOST_WEBHOOK_URL` environment variable with the Mattermost webhook URL) somewhere and redirect sentry webhook on it with route name defined as Mattermost channel for each projects.
![sentry-webhook-integration-setup](docs/assets/sentry-webhook-integration-setup.png)

Then you setup [sentry issue alerts](https://docs.sentry.io/product/alerts/) as you like.
![sentry-issue-alert-creation](docs/assets/sentry-issue-alert-creation.png)

## Deployment

Currently deployed on ps-ops1, and routed via https://sentry-ingest.ttiltd.com

To build and push an image:

    docker build . -t registry.ttiltd.net/sentry-mattermost-sidecar:1.0.1 && docker push registry.ttiltd.net/sentry-mattermost-sidecar:1.0.1

Deployment is managed via Salt. Update the version in the docker build command and salt and invoke the `docker.containers` state on ps-ops1 to deploy a new version.
