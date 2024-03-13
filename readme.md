# GitHub Handler <svg version="1.1" xmlns="http://www.w3.org/2000/svg" transform="translate(0, 20)" width="120" height="60" xml:space="preserve"><path fill="#00acd7" d="M22.3 24.7c-.1 0-.2-.1-.1-.2l.7-1c.1-.1.2-.2.4-.2h12.6c.1 0 .2.1.1.2l-.6.9c-.1.1-.2.2-.4.2l-12.7.1zM17 27.9c-.1 0-.2-.1-.1-.2l.7-1c.1-.1.2-.2.4-.2h16.1c.1 0 .2.1.2.2l-.3 1c0 .1-.2.2-.3.2H17zm8.5 3.3c-.1 0-.2-.1-.1-.2l.5-.9c.1-.1.2-.2.4-.2h7c.1 0 .2.1.2.2l-.1.8c0 .1-.1.2-.2.2l-7.7.1zM62.1 24l-5.9 1.5c-.5.1-.6.2-1-.4-.5-.6-.9-1-1.7-1.3-2.2-1.1-4.4-.8-6.4.5-2.4 1.5-3.6 3.8-3.6 6.7 0 2.8 2 5.1 4.8 5.5 2.4.3 4.4-.5 6-2.3.3-.4.6-.8 1-1.3h-6.8c-.7 0-.9-.5-.7-1.1.5-1.1 1.3-2.9 1.8-3.8.1-.2.4-.6.9-.6h12.8c-.1 1-.1 1.9-.2 2.9-.4 2.5-1.3 4.9-2.9 6.9-2.5 3.3-5.8 5.4-10 6-3.5.5-6.7-.2-9.5-2.3-2.6-2-4.1-4.6-4.5-7.8-.5-3.8.7-7.3 3-10.3 2.5-3.3 5.8-5.4 9.9-6.1 3.3-.6 6.5-.2 9.3 1.7 1.9 1.2 3.2 2.9 4.1 5 .1.4 0 .5-.4.6z"/><path fill="#00acd7" d="M73.7 43.5c-3.2-.1-6.1-1-8.6-3.1-2.1-1.8-3.4-4.1-3.8-6.8-.6-4 .5-7.5 2.9-10.6 2.6-3.4 5.7-5.1 9.9-5.9 3.6-.6 7-.3 10 1.8 2.8 1.9 4.5 4.5 5 7.9.6 4.8-.8 8.6-4 11.9-2.3 2.4-5.2 3.8-8.4 4.5-1.1.2-2.1.2-3 .3zm8.4-14.2c0-.5 0-.8-.1-1.2-.6-3.5-3.8-5.5-7.2-4.7-3.3.7-5.4 2.8-6.2 6.1-.6 2.7.7 5.5 3.2 6.7 1.9.8 3.9.7 5.7-.2 2.9-1.4 4.4-3.7 4.6-6.7z"/></svg>

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

Amir Farzamnia made GitHub Handler with ❤️ under the [MIT](LICENSE) license.