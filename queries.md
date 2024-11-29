# Queries

## How to run

Check the [README.md](README.md) file for instructions on how to run the project.  
After the seeder has finished its work, you can run the project without the `--profile seed` flag, because the database will already be populated.  
PgAdmin will be available at `http://localhost:5050`.  
You can run the queries from this file in the PgAdmin query tool.

## Simple

### Find all cards assigned to a specific user

```sql
SELECT cards.id, cards.title, cards.content, cards.due_date
FROM users
JOIN card_assignees ON users.id = card_assignees.user_id
JOIN cards ON card_assignees.card_id = cards.id
WHERE users.username = 'username';
```

### List all lists belonging to a specific board, ordered by their position

```sql
SELECT lists.id, lists.title, lists.order
FROM boards
JOIN lists ON boards.id = lists.board_id
WHERE boards.title = 'board_title'
ORDER BY lists.order;
```

### Find all comments made by a user on a specific card

```sql
SELECT comments.text, comments.created_at
FROM comments
JOIN users ON comments.user_id = users.id
JOIN cards ON comments.card_id = cards.id
WHERE users.username = 'username' AND cards.title = 'card_title';
```

### Retrieve all attachments for cards on a specific board

```sql
SELECT attachments.file_url
FROM attachments
JOIN cards ON attachments.card_id = cards.id
JOIN lists ON cards.list_id = lists.id
JOIN boards ON lists.board_id = boards.id
WHERE boards.title = 'board_title';
```

## Medium

### Find the top 3 most commented cards on a specific board

```sql
SELECT cards.title, COUNT(comments.id) AS comment_count
FROM cards
JOIN comments ON cards.id = comments.card_id
JOIN lists ON cards.list_id = lists.id
JOIN boards ON lists.board_id = boards.id
WHERE boards.title = 'board_title'
GROUP BY cards.id
ORDER BY comment_count DESC
LIMIT 3;
```

### Sort board members by the number of comments they made on cards on a specific board

```sql
SELECT users.username, COUNT(comments.id) AS comment_count
FROM comments
JOIN users ON comments.user_id = users.id
JOIN cards ON comments.card_id = cards.id
JOIN lists ON cards.list_id = lists.id
JOIN boards ON lists.board_id = boards.id
WHERE boards.title = 'board_title'
GROUP BY users.id
ORDER BY comment_count DESC;
```
