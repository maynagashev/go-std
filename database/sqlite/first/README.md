## Открытие базы данных

```bash
./sqlite3 video.db
```

## Создание таблицы

```sql
CREATE TABLE videos(
    "video_id" TEXT,
    "trending_date" TEXT,
    "title" TEXT,
    "channel_title" TEXT,
    "category_id" INTEGER,
    "publish_time" TEXT,
    "tags" TEXT,
    "views" INTEGER,
    "likes" INTEGER,
    "dislikes" INTEGER,
    "comment_count" INTEGER,
    "thumbnail_link" TEXT,
    "comments_disabled" BOOLEAN,
    "ratings_disabled" BOOLEAN,
    "video_error_or_removed" BOOLEAN,
    "description" TEXT
); 
```

## Наполение таблицы

```sqlite
.mode csv videos
.import USvideos.csv videos 
```

Проверка
```sqlite
.mode table
SELECT video_id,title, views from videos ORDER BY views LIMIT 10;
```