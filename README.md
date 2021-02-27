## CodelyTV - Go HTTP API - Hexagonal Architecture

This repository contains the code examples used on the .

### Requirements

- Go v1.15+
- MySQL (see below).

### Contents

This project has been designed as a single Go module with multiple applications.
Each folder contains a completely functional application (binary) that can be executed isolated.

Each folder corresponds to one of the lessons / videos:
1. [`08-03-debugging`](./08-03-debugging) - Generando la imagen de Docker y depurando errores

### Usage

To execute the application from any lesson, just run:

```sh
export COURSE_LESSON=02-04-domain-validations; go run $COURSE_LESSON/cmd/api/main.go 
```

Replacing `COURSE_LESSON` value by any of the available ones.

#### Simple examples

Some lessons only contain a single `main.go` file with a few lines of code.
To run one of those lessons, just run:

```sh
export COURSE_LESSON=01-01-your-first-http-endpoint; go run $COURSE_LESSON/main.go 
```

#### MySQL & Docker

From `02-01-post-course-endpoint` on, the application on each directory relies
on a MySQL database. So, to simplify its execution, we've added a
`docker-compose.yaml` file with a MySQL container already set up.

To run it, just execute:

```sh
docker-compose up -d 
```

You can also use your own MySQL instance. Note that those applications
expects a MySQL instance to be available on `localhost:3306`,
identified by `codely:codely` and with a `codely` database.

To set up your database, you can execute the `schema.sql` file
present on the `sql` directory. It's automatically loaded if
you use the provided `docker-compose.yaml` file.

#### Tests

To execute all tests, just run:

```sh
go test ./... 
```

To execute only the tests present in one of the lessons, run:

```sh
go test ./02-04-domain-validations/... 
```