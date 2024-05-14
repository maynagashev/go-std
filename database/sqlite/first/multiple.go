package first

import (
	"context"
	"database/sql"
)

func QueryVideos(ctx context.Context, db *sql.DB, limit int) ([]Video, error) {
	videos := make([]Video, 0, limit)

	rows, err := db.QueryContext(ctx,
		"SELECT video_id, title, views from videos ORDER BY views LIMIT ?", limit)
	if err != nil {
		return nil, err
	}

	// обязательно закрываем перед возвратом функции
	defer rows.Close()

	// пробегаем по всем записям
	for rows.Next() {
		var v Video
		err = rows.Scan(&v.Id, &v.Title, &v.Views)
		if err != nil {
			return nil, err
		}

		videos = append(videos, v)
	}

	// проверяем на ошибки
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return videos, nil
}

type Channel struct {
	Title    string
	Views    int64
	Videos   int64
	AvgViews float64
}

// QueryPopularChannels 30 самых популярных каналов, отсортированных по убыванию количества просмотров.
// Используйте GROUP BY и агрегирующие функции SUM(), COUNT(), AVG().
func QueryPopularChannels(ctx context.Context, db *sql.DB, limit int) ([]Channel, error) {
	channels := make([]Channel, 0, limit)

	rows, err := db.QueryContext(ctx,
		"SELECT channel_title, SUM(views) as views, COUNT(video_id) as videos, AVG(views) as avg_views "+
			"FROM videos GROUP BY channel_title "+
			"ORDER BY views "+
			"DESC LIMIT ?", limit)
	if err != nil {
		return nil, err
	}

	// обязательно закрываем перед возвратом функции
	defer rows.Close()

	// пробегаем по всем записям
	for rows.Next() {
		var c Channel
		err = rows.Scan(&c.Title, &c.Views, &c.Videos, &c.AvgViews)
		if err != nil {
			return nil, err
		}

		channels = append(channels, c)
	}

	// проверяем на ошибки
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return channels, nil
}
