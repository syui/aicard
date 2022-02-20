heroku api server

```sh
$ export DATABASE_URL=`heroku config:get DATABASE_URL`
$ export PORT=5000
$ curl localhost:5000/users/1
```
