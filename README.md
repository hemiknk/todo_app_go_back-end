# todo_app_go_back-end

Simple todo application written in golang

# Requirements
One of two databases can be used: SQLite - for development OR Postgresql for production
* go version 1.20 or hire
  * [GoLang](https://go.dev/)
* SQLite3 version 3.43.0 (development mode)
    * [SQLite3](https://www.sqlite.org/index.html)
* Postgresql version 16 (production mode)
    * [Postgresql](https://www.postgresql.org/) 


# ENV variables
| Name           | Example   | Description                                                                                                            |
|----------------|-----------|------------------------------------------------------------------------------------------------------------------------|
| DEV_ENV        | true      | Set true if you want to use DEV environment andSQLite3 database, false if you want to use PROD and PostgreSQL database |
| SQLITE_DB_PATH | ./todo.db | path to SQLite3 database                                                                                               |
| PG_DB_HOST     | localhost | PostgreSQL database host                                                                                               |
| PG_DB_PORT     | 5432      | PostgreSQL database port                                                                                               |
| PG_DB_PASSWORD | password  | PostgreSQL database password                                                                                           |
| PG_DB_NAME     | todo      | PostgreSQL database name                                                                                               |

# Run tests
* run tests `go test ./...`

# Local development
* clone repository `git clone git@bitbucket.org:hemikNK/todo_app_go_back-end.git`
* copy .env-sample to .env end set values
* run database migrations `go run cmd/migration/main.go up`
* run application `go run cmd/todoapp/main.go`
    * visit http://localhost:8018 in your browser

# Run in production
* copy .env-sample to .env end set values
* build application `go build -o todoapp cmd/todoapp/main.go`
* run application `./todoapp`

# Run in docker
* copy .env-sample to .env end set values
* docker build -t my-todo-app .
* docker run -p 8018:8018 my-todo-app
* visit http://localhost:8018 in your browser

# Run with docker-compose
* copy .env-sample to .env end set values
* docker-compose up
* visit http://localhost:8018 in your browser