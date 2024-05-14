package main

import (
	"context"
	"database/sql"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	_ "modernc.org/sqlite"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	// открываем соединение с БД
	db, err := sql.Open("sqlite", "newvideo.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	ctx := context.Background()
	_, err = db.ExecContext(ctx, `CREATE TABLE IF NOT EXISTS videos (
        "video_id" TEXT,
        "title" TEXT,
        "publish_time" TEXT,
        "tags" TEXT,
        "views" INTEGER
      )`)
	if err != nil {
		log.Fatal(err)
	}

	start := time.Now()
	// читаем записи из файла в слайс []Video вспомогательной функцией
	videos, err := readVideoCSV("USvideos.csv")
	if err != nil {
		log.Fatal(err)
	}
	// записываем []Video в базу данных
	// тоже вспомогательной функцией
	err = insertVideosTx(context.Background(), db, videos)
	if err != nil {
		log.Fatal(err)
	}
	// выводим время выполнения
	fmt.Println(time.Since(start))
	fmt.Printf("Всего csv-записей %v\n", len(videos))
}

type Video struct {
	Id          string    // video_id
	Title       string    // title
	PublishTime time.Time // publish_time
	Tags        []string  // tags
	Views       int       // views
}

// Вставка записей без использования транзакций (сильно медленнее)
func insertVideos(ctx context.Context, db *sql.DB, videos []Video) error {
	for _, v := range videos {
		// в этом случае возвращаемое значение не несёт полезной информации,
		// поэтому его можно игнорировать
		_, err := db.ExecContext(ctx,
			"INSERT INTO videos (video_id, title, publish_time, tags, views)"+
				" VALUES(?,?,?,?,?)", v.Id, v.Title, v.PublishTime,
			strings.Join(v.Tags, `|`), v.Views)
		if err != nil {
			return err
		}
	}
	return nil
}

// Вставка видео с использованием транзакций (быстрее примерно в 40 раз)
func insertVideosTx(ctx context.Context, db *sql.DB, videos []Video) error {
	// начинаем транзакцию
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	for _, v := range videos {
		// все изменения записываются в транзакцию
		_, err := tx.ExecContext(ctx,
			"INSERT INTO videos (video_id, title, publish_time, tags, views)"+
				" VALUES(?,?,?,?,?)", v.Id, v.Title, v.PublishTime,
			strings.Join(v.Tags, `|`), v.Views)
		if err != nil {
			// если ошибка, то откатываем изменения
			tx.Rollback()
			return err
		}
	}
	// завершаем транзакцию
	return tx.Commit()
}

func readVideoCSV(csvFile string) ([]Video, error) {
	// открываем csv файл
	file, err := os.Open(csvFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var videos []Video

	// определим индексы нужных полей
	const (
		Id          = 0 // video_id
		Title       = 2 // title
		PublishTime = 5 // publish_time
		Tags        = 6 // tags
		Views       = 7 // views
	)

	// конструируем Reader из пакета encoding/csv
	// он умеет читать строки csv-файла
	r := csv.NewReader(file)
	// пропустим первую строку с именами полей
	if _, err := r.Read(); err != nil {
		return nil, err
	}

	for {
		// csv.Reader за одну операцию Read() считывает одну csv-запись
		l, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		// инициализируем целевую структуру,
		// в которую будем делать разбор csv-записи
		v := Video{
			Id:    l[Id],
			Title: l[Title],
		}
		// парсинг строковых записей в типизированные поля структуры
		if v.PublishTime, err = time.Parse(time.RFC3339, l[PublishTime]); err != nil {
			return nil, err
		}
		tags := strings.Split(l[Tags], "|")
		for i, v := range tags {
			tags[i] = strings.Trim(v, `"`)
		}
		v.Tags = tags
		if v.Views, err = strconv.Atoi(l[Views]); err != nil {
			return nil, err
		}
		// добавляем полученную структуру в слайс
		videos = append(videos, v)
	}
	return videos, nil
}
