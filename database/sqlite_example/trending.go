package main

import (
	"database/sql"
	"time"
)

type Trend struct {
	T     time.Time
	Count int
}

func TrendingCount(db *sql.DB) ([]Trend, error) {
	rows, err := db.Query("SELECT trending_date, COUNT(trending_date) FROM videos" +
		" GROUP BY trending_date")
	if err != nil {
		return nil, err
	}
	trends := make([]Trend, 0)
	date := new(string)
	for rows.Next() {
		trend := Trend{}
		err = rows.Scan(date, &trend.Count)
		if err != nil {
			return nil, err
		}
		var t time.Time
		if t, err = time.Parse("06.02.01", *date); err != nil {
			return nil, err
		}
		trend.T = t
		trends = append(trends, trend)
	}
	return trends, nil
}
