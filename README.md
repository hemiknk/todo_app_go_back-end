# todo_app_go_back-end

Simple todo application written in golang

# Requironments
* SQLite3
    * [SQLite3](https://www.sqlite.org/index.html)
* go version 1.20
    * [GoLang](https://go.dev/)



# ENV variables
| Name | Example | Description |
| --- | --- | --- |
| HOST | localhost | Host where app will be served |
| PORT | 8018 | Port where app will be served |
| DEV_ENV | true | Set true if you want to use DEV environment andSQLite3 database, false if you want to use PROD and PostgreSQL database |
| SQLITE_DB_PATH | ./todo.db | path to SQLite3 database |
| PG_DB_HOST | localhost | PostgreSQL database host |
| PG_DB_PORT | 5432 | PostgreSQL database port |
| PG_DB_PASSWORD | password | PostgreSQL database password |
| PG_DB_NAME | todo | PostgreSQL database name |

# Run tests
* run tests `go test ./...`

# Local development
* clone repository `git clone git@bitbucket.org:hemikNK/todo_app_go_back-end.git`
* copy .env-sample to .env end set values
* run database migrations `go run cmd/migration/main.go up`
* run application `go run cmd/todoapp/main.go`
    * wisit http://localhost:8018 in your browser

# Run in production
* copy .env-sample to .env end set values
* build application `go build -o todoapp cmd/todoapp/main.go`
* run application `./todoapp`