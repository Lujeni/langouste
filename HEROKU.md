## How to
To install this project using Heroku, you will need:

1. A Heroku account, available for free from [Heroku.com](http://heroku.com)
2. A Heroku CLI, available for free from [Heroku.com](https://devcenter.heroku.com/articles/heroku-cli)
3. [**common part**](COMMON.md) will be necessary for a step.

## Deployment to Heroku
```bash
    $ cd <path_to_langouste_repository>
    
    $ heroku login
    
    $ heroku create

    $ git push heroku master

    # go to common part, check below

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
