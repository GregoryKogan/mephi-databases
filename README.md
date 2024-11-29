# mephi-databases

Laboratory work on the database course at MEPhI.  
Topic: Trello

This project is built in Go, uses [GORM](https://gorm.io/) as an ORM for PostgreSQL and [gofakeit](https://github.com/brianvoe/gofakeit) package to seed the database.

**Checkout the [queries.md](queries.md) file for possible queries to run.**

<details>
  <summary>Entity-relationship diagram</summary>
  <p align="center">
    <img src=https://github.com/user-attachments/assets/a238f8bc-5cea-41d4-b81b-e7b5e1c949cd width=50% />
  </p>
</details>

## Run

The project is dockerized. To run it, you need to have Docker installed on your machine.  
If you want to run seeder, you need to pass `--profile seed` to the `docker-compose` command.

```bash
docker compose --profile seed up --build
```

After the seeder has finished its work, you can run the project without the `--profile seed` flag, because the database will already be populated.

```bash
docker-compose up --build
```

### Configuration

Projects's configuration is stored in the `config.yml` file.  
Default configuration:

```yaml
seeder:
  load_batch_size: 2000
  create_batch_size: 5000
  entities: # Number of entities to create
    users: 10000
    # Average ratios of entities
    # Notice that labels and board members are the most numerous entities (30 times more than users)
    # Overall, the number of entities is 161 times more than the number of users
    # For example, if you have 10.000 users, you will have 1.610.000 entities in total
    boards_per_user: 3 # 3x users
    lists_per_board: 3 # 9x users
    cards_per_list: 2 # 18x users
    labels_per_board: 10 # 30x users
    comments_per_card: 0.5 # 9x users
    attachments_per_card: 0.333 # 6x users
    board_members_per_board: 10 # 30x users
    card_labels_per_card: 1.5 # 27x users
    card_assignees_per_card: 1.5 # 27x users
```

### Benchmark

On my machine (M1Pro MBP 14"), the seeder creates 10k users and 16.1M entities in roughly 40 seconds (with the default configuration).  
`pgdata` volume is about 2.7GB.

## PgAdmin

PgAdmin will be available at `http://localhost:5050`.

Credentials to connect to the database:

```plaintext
Host name/address: database
Port: 5432
Username: mephi
Password: mephi
```

<p align="center">
  <img src=https://github.com/user-attachments/assets/a5863b47-0bed-446a-bddd-651acd1dd367 width=50% />
</p>
