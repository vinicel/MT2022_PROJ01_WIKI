# Description

This is a school project. The requirements are to create 
a restfull api, using [Go](https://golang.org/), to 
manage some articles with comments. We must 
use JWT for the authentication, MVC pattern for the 
architecture, and [GORM](https://gorm.io/index.html) 
with MySQL for the database.

A complete documentation with a clean repository is mandatory.

## Table of Contents

- [Directory Structure](#directory-structure)
- [Model Structure](#model-structure)
- [Used Package](#used-package)
- [Api Documentation](#api-documentation)
- [Installation](#installation)
  
## Directory Structure

``` bash
<Wiki-Go>/
├─ src/
│   ├─ Controllers/
│   ├─ Models/
│   ├─ Services/
│   ├─ Views/
├─ main.go
└─ README.md
```

## Model Structure

Todo => add database schema
https://user-images.githubusercontent.com/29546258/107534201-fea3b800-6bbf-11eb-9258-7487036ba847.jpg

## Used Package

* [crypto](https://github.com/golang/crypto) - Go cryptography libraries (bcrypt)
* [mux](https://github.com/gorilla/mux) - HTTP request router and dispatcher
* [gorm](https://gorm.io/) - ORM library for Go
* [jwt](https://github.com/dgrijalva/jwt-go) - Go implementation of [JSON Web Tokens](https://self-issued.info/docs/draft-ietf-oauth-json-web-token.html)
* [swagger](https://github.com/go-swagger/go-swagger) - Go implementation of [Swagger 2.0](https://github.com/OAI/OpenAPI-Specification/blob/master/versions/2.0.md)

## Api Documentation

Todo => add link to swagger documentation

## Installation

* Build the project
``` bash
go build main.go
```
* todo => add more installation steps
``` bash
(I don't know Go)
```

