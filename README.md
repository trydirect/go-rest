# go-rest

RESTful API development template based on GIN framework for development purpose

## Stack includes
* Go 1.13
* Rest API
* Swagger
* GIN 1.5
* Alpine 3.11

## Note
Before installing this project, please, make sure you have installed docker and docker-compose

To install docker execute: 
```sh
$ curl -fsSL https://get.docker.com -o get-docker.sh
$ sh get-docker.sh
$ pip install docker-compose
```

## Installation
Clone this project into your work directory:
```sh
$ git clone https://github.com/trydirect/go-rest.git
```
Then build it via docker-compose:
```sh
$ cd go-rest
$ docker-compose up -d
or
$ docker-compose -f docker-compose-debug.yml up
```

Now, let's check the result
```
$ curl http://localhost:8080/api/v1/accounts
[{"id":1,"name":"account_1","uuid":"00000000-0000-0000-0000-000000000000"},{"id":2,"name":"account_2","uuid":"00000000-0000-0000-0000-000000000000"},{"id":3,"name":"account_3","uuid":"00000000-0000-0000-0000-000000000000"}]
```

Go to page http://localhost:8080/swagger/index.html to see generated documentation.

## Generating OpenAPI

1. Add comments to your API source code, See [Declarative Comments Format](https://github.com/swaggo/swag/blob/master/README.md#declarative-comments-format).

2. Download swag by using:
```sh
$ go get -u github.com/swaggo/swag/cmd/swag
```
Or download a pre-compiled binary from the [release page](https://github.com/swaggo/swag/releases).

3. Run `swag init` in the project's root folder. This will parse your comments and generate the required files (`docs` folder and `docs/docs.go`).
```sh
$ swag init -g app/main.go
```

# Contributing

1. Fork it (https://github.com/trydirect/go-rest/fork)
2. Create your feature branch (git checkout -b feature/fooBar)
3. Commit your changes (git commit -am 'Add some fooBar')
4. Push to the branch (git push origin feature/fooBar)
5. Create a new Pull Request


# Support Development

[![Donate](https://img.shields.io/badge/Donate-PayPal-green.svg)](https://www.paypal.com/cgi-bin/webscr?cmd=_s-xclick&hosted_button_id=2BH8ED2AUU2RL)