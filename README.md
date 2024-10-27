# mephi-databases

Laboratory work on the database course at MEPhI.  
Topic: Trello

This project is built in Go and uses [GORM](https://gorm.io/) as an ORM for PostgreSQL.

## Run

The project is dockerized. To run it, you need to have Docker installed on your machine.

```bash
docker-compose up --build
```

## PgAdmin

PgAdmin will be available at `http://localhost:5050`.

Credentials to connect to the database:

```
Host name/address: database
Port: 5432
Username: mephi
Password: mephi
```
