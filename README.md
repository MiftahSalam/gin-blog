# Simple Gin Blog

A small and simple API blog developed with go.

Developed to comply API spec from the [RealWorld](https://github.com/gothinkster/realworld)  
   
# Feature
- CRUD
- Authentication (JWT)
- Routing
- Pagination
- Input Validation
- Testing

# Demo
I deploy  this app on heroku. [Check Demo](https://gin-blog.herokuapp.com/api/) 

# Stacks
- Golang/Gin
- GORM (Golang ORM)
- Postgresql
- JWT-GO
- godotenv


# Getting started

## Install Golang

Make sure you have Go 1.13 or higher installed.

https://golang.org/doc/install

## Golang Environment Config

Set-up the standard Go environment variables according to latest guidance (see https://golang.org/doc/install#install).


## Install Dependencies
From the project root, run:
```
go build ./...
go test ./...
go mod tidy
```

## Application Environment Config
- Make sure database ready (installed and configured)
- create .env file in projectroot
- create and fill this env var in .env file
```
JWT_SECRET=
JWT_EXPIRED_IN=

DATABASE_HOST=
DATABASE_PORT=
DATABASE_USERNAME=
DATABASE_PASSWORD=
DATABASE_NAME=
DATABASE_LOGGING=
```

## Testing
Testing only available for every package (model, service, router). No root test provided. Model and service  are unit and integration testing. Router test is e2e testing.

To perform test, just go to folder package and execute:
```
go test ./...
```
or
```
go test ./... -cover
```
or
```
go test -v ./... -cover
```
depending on whether you want to see test coverage and how verbose the output you want.

## Todo
- Create client apps (different project)
- Cleanner code
- Chat
- Code structure optimize (I think some place can use interface)
- More advance article filtering (combine tag, author, favorite)