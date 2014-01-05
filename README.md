# Revel Guestbook

This is sample guestbook app with revel, gorp.


## upload to heroku

```
$ heroku create -b https://github.com/robfig/heroku-buildpack-go-revel.git
$ heroku addons:add cleardb:ignite
$ git push heroku master
$ heroku open
```