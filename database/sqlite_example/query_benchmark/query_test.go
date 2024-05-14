// query_test.go
package main

import (
	"database/sql"
	"log"
	"testing"

	_ "modernc.org/sqlite"
)

func BenchmarkQuery(b *testing.B) {
	db, err := sql.Open("sqlite", "newvideo.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		db.QueryRow("SELECT title, views FROM videos WHERE video_id = ?", "R39-E3uG5J0")
	}
}
