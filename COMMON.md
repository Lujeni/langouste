# How to
1. generate google your ID and secret, follow the [step](https://developers.google.com/identity/sign-in/web/devconsole-project)
  1. `Authorized JavaScript origins` field dos not require
  2. In the field `Authorized redirect URI` add the `domain` + `callback` endpoint, **http** and **https** versions.
    1. e.g. with heroku http(s)://nameless-river-68268.herokuapp.com/callback
    2. e.g. http(s)://<your_domain>/callback
