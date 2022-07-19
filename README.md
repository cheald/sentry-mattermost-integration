# sentry-mattermost-integration

This is a fork of [sentry-mattermost-sidecar](https://github.com/Its-Alex/sentry-mattermost-sidecar) designed to work with Sentry's newer internal integrations rather than its legacy webhooks.

Events handled

- [x] Bug alert triggered
- [ ] Issue created
- [ ] Issue resolved

## How to use

1. Create a [Mattermost incoming webhook](https://docs.mattermost.com/developer/webhooks-incoming.html) integration.

    ![mattermost-incoming-webhook-integration-setup](docs/assets/mattermost-incoming-webhook-integration-setup.png)

2. Deploy the [docker image](https://hub.docker.com/r/itsalex/sentry-mattermost-sidecar) and specify your Mattermost web hook URL in the `SMS_MATTERMOST_WEBHOOK_URL` environment variable

3. Create a Sentry integration which points to your deployment, including a channel slug to post to.

    ![sentry-webhook-integration-setup](docs/assets/sentry-integration-setup.png)

4. Set up [sentry issue alerts](https://docs.sentry.io/product/alerts/) as you like.

    ![sentry-issue-alert-creation](docs/assets/sentry-issue-alert-creation.png)

## Development

### Requirements

- `docker`
- `go`
- `bash`

To start you must launch dev environment:

```sh
$ ./scripts/up.sh
```

This will launch images in [`docker-compose.yml`](./docker-compose.yml).

An image named `workspace` with golang is used as a isolated container to develop. You can use [`enter-workspace.sh`](./scripts/enter-workspace.sh) to enter inside it:

```sh
$ ./scripts/enter-workspace.sh
```

You can build with:

```sh
$ ./scripts/build.sh
```

You can test an example sentry webhook with:

```sh
$ ./scripts/test-request.sh
```

Then you can see the converted request that will be send to mattermost using:

```sh
$ ./scripts/get-last-request-result.sh
```

## Deploy

This image is automatically deployed and versionned as a docker image at [itsalex/sentry-mattermost-sidecar](https://hub.docker.com/r/itsalex/sentry-mattermost-sidecar).

To deploy a new tag use [`./scripts/create-and-push-tag.sh`](scripts/create-and-push-tag.sh):

```sh
$ ./scripts/create-and-push-tag.sh 1.0.0
```
