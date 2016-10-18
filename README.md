# Introduction

Langouste is a web apps designed for working with Google Calendar and Mattermost.

Is inspired by the official slack command.

Just type `/hangout` in any channel or direct message to create unique Hangout link.

![langouste screenshot](https://github.com/lujeni/langouste/raw/master/langouste.png)

# Usage
## Requirements
  * Download [Latest Version](https://github.com/lujeni/langouste/releases/latest)
  * Follow the `step 1` of the [Google calendar quickstart](https://developers.google.com/google-apps/calendar/quickstart/go)


## Environnement
```bash
# export the necessary variables.
export langoustePort=8000
export ClientSecret="YksceBugZGagGi-18Nn"   # from the google quick start
export ClientID="4259064-r75o49rnct7mvuf08.apps.googleusercontent.com" # from the google quick start
```

## Setup
### Launch the web server 
```bash
./langouste
```

### Generate the token
```bash
visit the http://localhost:8000/
```

### Mattermost
* Create a `slash` command
** url: http://localhost:8000/
** method: POST
