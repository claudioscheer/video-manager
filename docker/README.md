## After starting the stack

- Create the database;
```sql
CREATE DATABASE IF NOT EXISTS video_manager;
```

- Create a user;
```sql
GRANT ALL PRIVILEGES ON video_manager.* TO 'videomanager'@'%' IDENTIFIED BY 'videomanager';
```