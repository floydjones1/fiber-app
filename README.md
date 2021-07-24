# Golang Web Server Demo (go-fiber) [![license](https://img.shields.io/github/license/DAVFoundation/captain-n3m0.svg?style=flat-square)](https://github.com/DAVFoundation/captain-n3m0/blob/master/LICENSE)
This app is to be used as a boiler plate for a golang web server using go-fiber to drive the major piece. Creation of routes and handlers should look familiar if you ever have ever made an express web server.

## Why I think this makes a really good starting point?
1. [go-fiber](https://github.com/gofiber/fiber) has amazing benchmarks for performance and has other supporting modules that are available for JWT auth, swagger docs, rate limiter and more!
2. [XORM](https://gobook.io/read/gitea.com/xorm/manual-en-US/) is a well balanced client which can be used as an ORM or raw queries when you need it.
3. XORM also has built in caching to help performance if needed.
4. [Goose](https://github.com/pressly/goose) allows to write migrations in both SQL and go. (others might to it aswell)
5. Follows Golang widely accepted folder structure [conventions.](https://github.com/golang-standards/project-layout)

## More to know
This app was made to be usable on both Windows and Mac. `make` isn't available by default on Windows but you can install it regardless

## Other pieces of tools used are:
```
go-fiber <--- Web Server
xorm <-- ORM/Database Client
goose <-- Database Migration
zeroLog <-- Logging
modd <-- Live Reload
docker <-- dev environemnt
postgres <-- Database
make <-- dev tool
```

## How to start App?

```
make tools <-- Downloads required go binaries
make up <-- Startup Postgres 
make start <-- Runs go-fiber
```