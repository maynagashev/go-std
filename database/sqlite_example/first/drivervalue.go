package first

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"errors"
	"fmt"
	"strings"
)

// Valuer приводит сложные типы и структуры к простому типу
type Valuer interface {
	Value() (driver.Value, error)
}

// Scanner приводит простой тип к сложным типам и структурам Go
type Scanner interface {
	Scan(src any) error
}

// Value — функция реализующая интерфейс driver.Valuer
func (tags Tags) Value() (driver.Value, error) {
	// преобразуем []string в string
	if len(tags) == 0 {
		return "", nil
	}
	return strings.Join(tags, "|"), nil
}

func (tags *Tags) Scan(value interface{}) error {
	// если `value` равен `nil`, будет возвращён пустой массив
	if value == nil {
		*tags = Tags{}
		return nil
	}

	sv, err := driver.String.ConvertValue(value)
	if err != nil {
		return fmt.Errorf("cannot scan value. %w", err)
	}

	v, ok := sv.(string)
	if !ok {
		return errors.New("cannot scan value. cannot convert value to string")
	}
	*tags = strings.Split(v, "|")

	// удаляем кавычки у тегов
	for i, v := range *tags {
		(*tags)[i] = strings.Trim(v, `"`)
	}
	return nil
}

func QueryTagVideos(ctx context.Context, db *sql.DB, limit int) ([]Video, error) {
	videos := make([]Video, 0, limit)

	rows, err := db.QueryContext(ctx, "SELECT video_id, title, tags from videos "+
		"GROUP BY video_id ORDER BY views LIMIT ?", limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var v Video
		// все теги должны автоматически преобразоваться в слайс v.Tags
		err = rows.Scan(&v.Id, &v.Title, &v.Tags)
		if err != nil {
			return nil, err
		}
		videos = append(videos, v)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return videos, nil
}
