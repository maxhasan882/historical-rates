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
> i.   GET: localhost:port/rates/latest   
> ii.  GET: localhost:port/rates/YYYY-MM-DD   
> iii. GET: localhost:port/rates/analyze

### ✍️ Bonus
> Use only the standard library. It is OK if a library is needed for TDD or database connection.   
>>No library other than the standard library was used.
I have used the following 3 external libraries for DB connection and testing.
> >>github.com/golang/mock v1.6.0 (for mocking)   
github.com/lib/pq v1.10.6 (for db connection)   
github.com/stretchr/testify v1.8.0 (for testing)
# 📋 Folder Structure
```
historical-rates
    ├── build
    │   ├── init.sql
    │   ├── docker-compose.yml
    │   ├── Dockerfile
    ├── cmd
    │   ├── app
    │   │   └── main.go                        # application entry point
    │   └── env                                # custom env loader
    │       ├── loader.go
    │       └── loader_test.go
    ├── internal                               # main source directory
    │   └── app
    │       ├── adapter                        # outer layer. all framework, external database and middlewares related code 
    │       │   ├── controller
    │       │   │   ├── rate.go
    │       │   │   ├── rate_test.go
    │       │   │   ├── response_handler.go    # common respponse handler
    │       │   │   ├── response_handler_test.go
    │       │   │   └── server.go
    │       │   ├── db
    │       │   │   └── connections
    │       │   │       └── pg_connection.go   # postgres db connection
    │       │   ├── repository                 # repository implementation
    │       │   │   ├── loader.go
    │       │   │   └── rate.go
    │       │   ├── route.go                  # custom route parser
    │       │   ├── route_test.go
    │       │   └── utils                     # utils functions
    │       │       ├── common.go
    │       │       ├── common_test.go
    │       │       ├── remote.go
    │       │       └── remote_test.go
    │       ├── application                   # middle layer. mainly deals with business logic
    │       │   └── usecase
    │       │       ├── loader.go
    │       │       ├── loader_test.go
    │       │       ├── rate.go
    │       │       └── rate_test.go
    │       └── domain                       # inner layer. all schema and repository defination
    │           ├── rate.go
    │           └── repository
    │               ├── loader.go
    │               ├── mocks                # interface moc for testing
    │               │   ├── loader_mock.go
    │               │   └── rate_mock.go
    │               └── rate.go
    ├── go.mod
    ├── go.sum
    ├── README.md
    └── .env
```
# ❓ Challenges and Solution
* Custom route management
  * As I'm not using any library other than standard library, so routing management was a bit challenging. That's why I have decided to write a custom simple route parser which can solve the purpose of the task. For parsing route I have followed tree mechanism. 
* Load .env 
  * There is no inbuilt .env file loader in golang so, I have written a custom .env loader for reading variables from a file and set it to environment.
