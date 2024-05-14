package main

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	_ "modernc.org/sqlite"
)

// Video — структура видео.
type Video struct {
	Id    string
	Title string
	Views int64
	Tags  Tags
}

type Tags []string

func main() {
	db, err := sql.Open("sqlite", "db/video.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	if err = db.PingContext(ctx); err != nil {
		panic(err)
	}

	countVideos(db)

	selectRow(db)

	desc, _ := getDesc(context.Background(), db, "0EbFotkXOiA")
	fmt.Println("description:", desc)

	// Самый популярный ролик
	mostPopular(db)

	// Простая выборка записей
	videos, err := QueryVideos(context.Background(), db, 20)
	if err != nil {
		return
	}
	fmt.Printf("Videos: %v\n\n", videos)

	// Выборка с агрегацией
	channels, err := QueryPopularChannels(context.Background(), db, 30)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	fmt.Printf("Channels: %#v\n", channels)

	// Парсинг кастомного типа tags
	list, err := QueryTagVideos(context.Background(), db, 5)
	if err != nil {
		panic(err)
	}
	// для теста проверим, какие строки содержит v.Tags
	// выведем по 4 первых тега
	for _, v := range list {
		length := 4
		if len(v.Tags) < length {
			length = len(v.Tags)
		}
		fmt.Println(strings.Join(v.Tags[:length], " # "))
	}

	// Кол-во роликов по каждому дню
	days, err := TrendingCount(db)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Trending by days: %v\n", days)
}

// title (название), views (количество просмотров) и channel_title (название канала)
// самого популярного ролика на YouTube.
func mostPopular(db *sql.DB) {
	row := db.QueryRowContext(context.Background(),
		"SELECT title, views, channel_title "+
			"FROM videos ORDER BY views DESC LIMIT 1")
	var (
		title  string
		views  int
		chname string
	)
	err := row.Scan(&title, &views, &chname)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Most popular: %s | %d | %s \r\n", title, views, chname)
}

func selectRow(db *sql.DB) {
	row := db.QueryRowContext(context.Background(),
		"SELECT title, likes, comments_disabled "+
			"FROM videos ORDER BY likes DESC LIMIT 1")
	var (
		title  string
		likes  int
		comdis bool
	)
	// порядок переменных должен соответствовать порядку колонок в запросе
	err := row.Scan(&title, &likes, &comdis)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s | %d | %t \r\n", title, likes, comdis)
}

func countVideos(db *sql.DB) {

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	// не забываем освободить ресурс
	defer cancel()

	// делаем запрос
	row := db.QueryRowContext(ctx, "SELECT COUNT(*) as count FROM videos")
	// готовим переменную для чтения результата
	var id int64
	err := row.Scan(&id) // разбираем результат
	if err != nil {
		panic(err)
	}
	fmt.Println(id)
}

// Проверка парсинга null string
func getDesc(ctx context.Context, db *sql.DB, id string) (string, error) {
	row := db.QueryRowContext(ctx,
		"SELECT description FROM videos WHERE video_id = ?", id)
	var desc sql.NullString

	err := row.Scan(&desc)
	if err != nil {
		return "", err
	}
	if desc.Valid {
		return desc.String, nil
	}
	return "-----", nil
}
