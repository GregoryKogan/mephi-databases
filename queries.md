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

### List all cards with due dates in the next 7 days, grouped by their list title on a specific board

```sql
SELECT lists.title AS list_title, cards.title AS card_title, cards.due_date
FROM cards
JOIN lists ON cards.list_id = lists.id
JOIN boards ON lists.board_id = boards.id
WHERE boards.title = 'board_title'
AND cards.due_date BETWEEN NOW() AND NOW() + INTERVAL '7 days'
ORDER BY lists.title, cards.due_date;
```

<!-- TODO: Implement created_at timestamp generation in the seeder -->

### Retrieve board members and their roles for a specific board ordered by their join date

```sql
SELECT users.username, board_roles.title AS role_title, board_members.created_at AS join_date
FROM board_members
JOIN users ON board_members.user_id = users.id
JOIN board_roles ON board_members.board_role_id = board_roles.id
WHERE board_members.board_id = (SELECT id FROM boards WHERE title = 'board_title')
ORDER BY join_date DESC;
```

## Complex

### Identify the most productive users across boards

Find users who contribute the most to boards based on the number of cards they are assigned, comments they make, and their membership status, combining these metrics into a productivity score.

```sql
WITH user_stats AS (
    SELECT
        users.id AS user_id,
        users.username,
        COUNT(DISTINCT card_assignees.card_id) AS assigned_cards_count,
        COUNT(DISTINCT comments.id) AS comments_count,
        COUNT(DISTINCT board_members.board_id) AS memberships_count
    FROM users
    LEFT JOIN card_assignees ON users.id = card_assignees.user_id
    LEFT JOIN comments ON users.id = comments.user_id
    LEFT JOIN board_members ON users.id = board_members.user_id
    GROUP BY users.id
)
SELECT
    username,
    assigned_cards_count,
    comments_count,
    memberships_count,
    (assigned_cards_count * 2 + comments_count + memberships_count * 1.5) AS productivity_score
FROM user_stats
ORDER BY productivity_score DESC
LIMIT 10;
```

### Track board activity trends

Calculate the daily average number of comments and card assignments on all boards, and track the 1-day/7-day/30-day moving averages.

```sql
WITH daily_activity AS (
    SELECT
        DATE(comments.created_at) AS activity_date,
        COUNT(DISTINCT comments.id) AS daily_comments,
        COUNT(DISTINCT card_assignees.card_id) AS daily_assignments
    FROM comments
    FULL JOIN card_assignees ON DATE(comments.created_at) = DATE(card_assignees.created_at)
    GROUP BY DATE(comments.created_at)
),
moving_averages AS (
    SELECT
        activity_date,
        daily_comments,
        daily_assignments,
        AVG(daily_comments) OVER (ORDER BY activity_date ROWS BETWEEN 1 PRECEDING AND CURRENT ROW) AS comments_1_day_avg,
        AVG(daily_comments) OVER (ORDER BY activity_date ROWS BETWEEN 6 PRECEDING AND CURRENT ROW) AS comments_7_day_avg,
        AVG(daily_comments) OVER (ORDER BY activity_date ROWS BETWEEN 29 PRECEDING AND CURRENT ROW) AS comments_30_day_avg,
        AVG(daily_assignments) OVER (ORDER BY activity_date ROWS BETWEEN 1 PRECEDING AND CURRENT ROW) AS assignments_1_day_avg,
        AVG(daily_assignments) OVER (ORDER BY activity_date ROWS BETWEEN 6 PRECEDING AND CURRENT ROW) AS assignments_7_day_avg,
        AVG(daily_assignments) OVER (ORDER BY activity_date ROWS BETWEEN 29 PRECEDING AND CURRENT ROW) AS assignments_30_day_avg
    FROM daily_activity
)
SELECT * FROM moving_averages ORDER BY activity_date;
```

<!-- TODO: Add created_at timestamp to card_assignees -->

### Detect boards with uneven work distribution

Identify boards where some users are disproportionately assigned cards compared to others.

```sql
WITH user_card_counts AS (
    SELECT
        boards.id AS board_id,
        users.id AS user_id,
        COUNT(cards.id) AS card_count
    FROM boards
    JOIN lists ON lists.board_id = boards.id
    JOIN cards ON cards.list_id = lists.id
    LEFT JOIN card_assignees ca ON ca.card_id = cards.id
    LEFT JOIN users ON users.id = ca.user_id
    GROUP BY boards.id, users.id
),
board_card_distribution AS (
    SELECT
        board_id,
        STDDEV(card_count) AS card_count_stddev,
        AVG(card_count) AS card_count_avg,
        COUNT(DISTINCT user_id) AS user_count
    FROM user_card_counts
    GROUP BY board_id
)
SELECT
    boards.title AS board_title,
    card_count_stddev,
    card_count_avg,
    user_count
FROM boards
JOIN board_card_distribution d ON boards.id = d.board_id
WHERE d.user_count > 1
AND d.card_count_stddev > (d.card_count_avg * 0.5)
ORDER BY d.card_count_stddev DESC;
```

### Calculate user retention metrics for boards

Determine how many users actively participate in boards over three periods: the last 7 days, last 30 days, and overall.

```sql
WITH user_activity AS (
    SELECT
        board_members.board_id,
        board_members.user_id,
        MAX(GREATEST(
            COALESCE(MAX(comments.created_at), '1970-01-01'),
            COALESCE(MAX(card_assignees.created_at), '1970-01-01')
        )) AS last_active_date
    FROM board_members
    LEFT JOIN comments ON board_members.user_id = comments.user_id
    LEFT JOIN card_assignees ON board_members.user_id = card_assignees.user_id
    GROUP BY board_members.board_id, board_members.user_id
)
SELECT
    board_id,
    COUNT(CASE WHEN last_active_date >= NOW() - INTERVAL '7 days' THEN 1 END) AS active_last_7_days,
    COUNT(CASE WHEN last_active_date >= NOW() - INTERVAL '30 days' THEN 1 END) AS active_last_30_days,
    COUNT(*) AS active_all_time
FROM user_activity
GROUP BY board_id;
```

<!-- TODO: Add created_at timestamp to card_assignees -->

### Find struggling members

Identify users who are frequently associated with overdue cards and the average overdue duration.

```sql
SELECT
    users.username,
    COUNT(cards.id) AS overdue_cards_count,
    AVG(EXTRACT(DAY FROM NOW() - cards.due_date)) AS avg_overdue_days
FROM users
JOIN card_assignees ON users.id = card_assignees.user_id
JOIN cards ON card_assignees.card_id = cards.id
WHERE cards.due_date < NOW()
GROUP BY users.id
HAVING COUNT(cards.id) > 3
ORDER BY avg_overdue_days DESC;
```

<!-- TODO: Implement due_date timestamp generation in the seeder -->

### Find "hot topics" cards with the most engagement

Identify cards with the most interactions (comments, attachments, and assignees).

```sql
WITH card_engagement AS (
    SELECT
        cards.title AS card_title,
        COUNT(DISTINCT comments.id) AS comment_count,
        COUNT(DISTINCT attachments.id) AS attachment_count,
        COUNT(DISTINCT card_assignees.user_id) AS assignee_count,
    FROM cards
    LEFT JOIN comments ON cards.id = comments.card_id
    LEFT JOIN attachments ON cards.id = attachments.card_id
    LEFT JOIN card_assignees ON cards.id = card_assignees.card_id
    GROUP BY cards.id
)
SELECT
    card_title,
    comment_count,
    attachment_count,
    assignee_count,
    (comment_count + attachment_count + assignee_count) AS engagement_score
FROM card_engagement
ORDER BY engagement_score DESC
LIMIT 5;
```

### Find average load for each day of the week

Calculate average load per day of the week over the last 6 months, considering comments and card assignments.

```sql
WITH daily_activity AS (
    SELECT
        DATE(COALESCE(comments.created_at, card_assignees.created_at)) AS activity_date,
        EXTRACT(DOW FROM DATE(COALESCE(comments.created_at, card_assignees.created_at))) AS day_of_week,
        COUNT(DISTINCT comments.id) AS comment_count,
        COUNT(DISTINCT card_assignees.id) AS card_assignment_count
    FROM comments
    FULL JOIN card_assignees
        ON comments.user_id = card_assignees.user_id
        AND DATE(comments.created_at) = DATE(card_assignees.created_at)
    WHERE COALESCE(comments.created_at, card_assignees.created_at) >= NOW() - INTERVAL '6 months'
    GROUP BY activity_date, day_of_week
),
total_daily_activity AS (
    SELECT
        day_of_week,
        (comment_count + card_assignment_count) AS total_activity
    FROM daily_activity
)
SELECT
    day_of_week,
    AVG(total_activity) AS avg_daily_load
FROM total_daily_activity
GROUP BY day_of_week
ORDER BY day_of_week;
```

<!-- TODO: Add created_at timestamp to card_assignees -->
