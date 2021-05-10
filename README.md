# Users API example

## Overview

Here you can see an example of the Users RESP API made in Golang using gorilla
mux as a router and postgres + gorm as a database. In addition, this project 
uses automatic generation of documentation using swagger (swaggo)

## How to run?

To run everything in docker-compose you can just write:

```bash
make start
```

And then call for example:

http://localhost:8080/user/0

To check API docs open:

http://localhost:8080/swagger/

To run tests with coverage, race condition and disabled cache, run:
```bash
make init
make test
```

### Project Navigation

- [CMD start file](./cmd/users/users.go)
- [Main logic](./internal/users.go)
- [Handlers](./internal/handler/user.go)
- [Docker-compose file](docker-compose.yml)

### ENV variables

You can use env variables for configuration:

`PORT` - port to handle requests

`DB_USER` - user for postgres

`DB_PASS` - password for postgres

`DB_HOST` - host for postgres

`DB_NAME` - database name for postgres

### What can I improve here?

I left a lot of **todo** in the code for us to work on

