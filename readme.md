heroku api server

```sh
$ mkdir -p ~/go/src
$ cd ~/go/src
$ git clone https://github.com/syui/aicard
$ cd aircard
# go get -u github.com/kardianos/govendor
# govendor init
# govendor fetch +out

$ git add .
$ git commit -m "heroku deploy"
$ heroku git:remote -a $APP
$ git push heroku main
```

```sh
$ curl --location --request POST "https://$APP.herokuapp.com/users" \
		--header 'Content-Type: application/json' \
		--data-raw '{
			"name": "syui"
		}'
$ curl --location --request PUT "https://$APP.herokuapp.com/users/0" \
		--header 'Content-Type: application/json' \
		--data-raw '{
			"name": "syui"
		}'
$ curl https://$APP.herokuapp.com/users
```


