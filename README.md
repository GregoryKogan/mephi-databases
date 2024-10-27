# mephi-databases

Laboratory work on the database course at MEPhI.  
Topic: Trello

This project is built in Go and uses [GORM](https://gorm.io/) as an ORM for PostgreSQL.

<details>
  <summary>Entity-relationship diagram</summary>
  <p align="center">
    <img src=https://github.com/user-attachments/assets/a238f8bc-5cea-41d4-b81b-e7b5e1c949cd width=50% />
  </p>
</details>

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
<p align="center">
  <img src=https://github.com/user-attachments/assets/a5863b47-0bed-446a-bddd-651acd1dd367 width=50% />
</p>
