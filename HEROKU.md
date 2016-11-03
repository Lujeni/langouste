## How to
To install this project using Heroku, you will need:

1. A Heroku account, available for free from [Heroku.com](http://heroku.com)
2. A Heroku CLI, available for free from [Heroku.com](https://devcenter.heroku.com/articles/heroku-cli)

## Deployment to Heroku
```bash
    $ cd <path_to_langouste_repository>
    
    $ heroku login
    
    $ heroku create

    $ git push heroku master

    # set the necessary Google stuff.
    $ heroku config:set ClientSecret=GNn5lhZWj1XlAv1876xc1

    $ heroku config:set ClientID=429064-6g0jf3h0fku7c4ct2nliceo4cn.apps.googleusercontent.com

    $ heroku open
```

## Setup
### Generate the token
```bash
go to http://<heroku_domain>/
```

### Mattermost
* Create the [Slash command](https://docs.mattermost.com/developer/slash-commands.html#set-up-a-custom-command)
1. url: http://<heroku_domain>/
2. request method: POST
