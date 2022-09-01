<br><br><br><br><br><br><br>
<br>
<br>
<h1 align="center">
  <img alt="cgapp logo" src="https://raw.githubusercontent.com/create-go-app/cli/master/.github/images/cgapp_logo%402x.png" width="224px"/><br/>
 Historical Exchange Rate Rest API Using Golang
</h1>
<p align="center">An API server using <b>backend</b> (Golang) and <b>database</b> (Postgres) containerised with (Docker)!</p>

<p align="center"><a href="#" 
target="_blank"><img src="https://img.shields.io/badge/Go-1.17+-00ADD8?style=for-the-badge&logo=go" alt="go version" /></a>&nbsp;<a href="#" target="_blank"><img src="https://img.shields.io/badge/-REST API-red?style=for-the-badge&logo=google-cloud&logoColor=white" alt="REST API" /></a>&nbsp;<a href="#" target="_blank"></a></p>

<p align="center"><img src="https://dl.circleci.com/status-badge/img/gh/maxhasan882/historical-rates/tree/master.svg?style=svg&circle-token=8c5074f9b541a302520bd3ea06b5e005315feed4" alt="test status" />
<img src="https://codecov.io/gh/maxhasan882/historical-rates/branch/circleci-project-setup/graph/badge.svg?token=0R4T7M7URO" alt="code coverage" /></p>
## 📖 Problem

Build an API server which load historical rate data from **ecb.europa.eu** and develop 3 endpoint
## ⚡️ Quick start

```shell
Run test 
     $ go test ./...  # run from root of the project directory
```
```shell
Using docker 
     $ docker-compose up --build --force-recreate # run it from root of the project directory
```
🔔 `Note`If you interested to run it from locally without Docker please ensure postgres database is up running also env is properly configured. Then create database table using **/build/init.sql** file.
- >go run /cmd/app/main.go
>Example .env file
>>_COMPOSE_FILE=build/docker-compose.yml
POSTGRES_USER=postgres
POSTGRES_PASSWORD=postgres
POSTGRES_HOST=database
POSTGRES_PORT=5432
DATABASE_NAME=historical_rate
SERVER_PORT=8080_

### ✍️ Task 1: Populate data
> It will automatically populate data at startup time.

### ✍️ Task 2, 3 and 4: Lets visit following endpoints
> i.   GET: localhost:port/rates/latest </br>
> ii.  GET: localhost:port/rates/YYYY-MM-DD </br>
> iii. GET: localhost:port/rates/analyze

### ✍️ Bonus
> Use only the standard library. It is OK if a library is needed for TDD or database connection. <br>
>>No library other than the standard library was used.
I have used the following 3 external libraries for DB connection and testing.
> >>github.com/golang/mock v1.6.0 (for mocking) <br>
github.com/lib/pq v1.10.6 (for db connection) <br>
github.com/stretchr/testify v1.8.0 (for testin)
# 📋 Folder Structure
```
book-info-graphql
├── build
│    ├── Dockerfile
│    ├── init.sql
├── cmd
│    ├── app
│    │    ├── config
│    │    │  └── loader.go
│    │    └── main.go
├── graph
│    ├── ***                  - All graph related code along with auto generated code
├── internal
│    └── app
│         ├── adapter         - Outer layer. All framework and external database and middlewares related code 
│         ├── application     - Middle layer. Usecase or buniness logic relaed code
│         │    └── usecase
│         └── domain          - Inner layer. Domain, interface and factory related code
│              ├── interface
│              └── factory
└── .env
```
# ❓ Challenges and Solution
* Custom route management
  * 
* Load .env 
  * 
