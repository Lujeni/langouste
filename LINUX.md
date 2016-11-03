## How to
To install this project, you will need:

1. `go` package, available with your package manager (e.g. apt-get install go)
2. `langouste` package, you can download it [Here](https://github.com/lujeni/langouste/releases/latest)

## Deployment
```bash
    # the default listen port is 8000
    # you can override it through the PORT variable
    $ export PORT=8888

    # set the necessary Google stuff.
    $ export ClientSecret="YksceBugZGagGi-18Nn"

    $ export ClientID="4259064-r75o49rnct7mvuf08.apps.googleusercontent.com"

    $ chmod 755 <path_to_langouste>/langouste

    $ ./<path_to_langouste>/langouste
```

## Setup
### Generate the token
```bash
visit http://<YOUR_DOMAIN>:<YOUR_PORT>/ with your browser
```

### Mattermost
* Create the [Slash command](https://docs.mattermost.com/developer/slash-commands.html#set-up-a-custom-command)
1. url: http://<YOUR_DOMAIN>:<YOUR_PORT>
2. request method: POST
