# GitHub Handler

GitHub Handler helps your apps manage events occurring within your GitHub repositories.

## Functionality

GitHub Handler uses GitHub's official webhooks. Each time a modification is made to your GitHub repository, GitHub dispatches a webhook request to the server where the GitHub Handler operates. The GitHub Handler then processes the incoming webhook and carries out your specified commands to manage alterations and events within your GitHub repositories.

## Configuring GitHub Webhook

- Go to your GitHub repo's "Settings" and then "Webhooks" and click on the "Add webhook" button.
- Use a custom payload URL for your webhook. This is the URL that the GitHub Handler is running on.
- You can select any content type for requests you want.
- Choose a secret key for your requests. This makes sure that all incoming requests are from GitHub and using your secret key.
- Select the events you wish to receive from GitHub, make sure that the "Active" checkbox is checked, and click on the "Add webhook" button.

## Installing GitHub Handler

- First, you need to install the Go programming language on your app server. [Click here to see how to install the Go language](https://go.dev/doc/install).
- Clone this repo on your app server:

```bash
git clone https://github.com/amirfarzamnia/github-handler
```

- Go to the GitHub Handler's directory by using:

```bash
cd github-handler
```

- Run the GitHub Handler using the below command:

```bash
go run handler.go -port "5050" -secret "my_secret_key" -command "./handler.sh"
```

That's it! GitHub Handler is now successfully installed and is ready to handle all incoming events.

## Configuration

Here are the configurations to run the `handler.go` file:

- `port`: server port to be listened on.
- `secret`: your secret key, which is used as the GitHub webhook's secret. Use an empty "" to not use secret validation.
- `command`: The command that will be executed after the request has been validated and successfully received. Use an empty "" to do nothing.

Here is the default configuration:

- `port`: 5050
- `secret`: my_secret_key
- `command`: ./handler.sh

## Notes

- It's recommended to use https payload urls such as `https://github.domain.com` and use SSL verification for more secure connections.
- Your secret key will be encoded in SHA-256 format by GitHub.
- GitHub Handler can also be executed without having a secret key, but if you do not use one, hackers may be able to send custom events with fake requests to your servers. Which can result in DDOS attacks.

---

Amir Farzamnia made GitHub Handler with ❤️ under the [MIT](./LICENSE) license.
