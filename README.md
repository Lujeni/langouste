# Introduction

Langouste is a web apps designed for working with Google Calendar and Mattermost.

Is inspired by the official slack command.

Just type `/hangout` in any channel or direct message to create unique Hangout link.

![langouste screen](https://github.com/Lujeni/langouste/blob/master/langouste.png)

# Usage
## Requirements
  * Download [Latest Version](https://github.com/lujeni/langouste/releases/latest)
  * Follow the `step 1` of the [Google calendar quickstart](https://developers.google.com/google-apps/calendar/quickstart/go)


## Environnement
```bash
# the default listen port is 8000
# you can override it through the PORT variable
export PORT=8888

# export the necessary variables.
export ClientSecret="YksceBugZGagGi-18Nn"   # from the google quick start
export ClientID="4259064-r75o49rnct7mvuf08.apps.googleusercontent.com" # from the google quick start
```

## Setup
### Launch the web server
```bash
# GET -> / generetate the temporary google token (oauth).
# POST -> / handle mattemost event from the slash command.
./langouste
```

### Generate the token
```bash
go to http://localhost:8000/
```

### Mattermost
* Create the [Slash command](https://docs.mattermost.com/developer/slash-commands.html#set-up-a-custom-command)
1. url: http://localhost:8000
2. request method: POST
