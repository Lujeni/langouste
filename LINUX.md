## How to
To install this project, you will need:

1. `go` package, available with your package manager (e.g. apt-get install go)
2. `langouste` package, you can download it [Here](https://github.com/lujeni/langouste/releases/latest)
3. [**common part**](COMMON.md) will be necessary for a step.

## Deployment
```bash
    # the default listen port is 8000
    # you can override it through the PORT variable
    $ export PORT=8888

    # go to common part, check below

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
