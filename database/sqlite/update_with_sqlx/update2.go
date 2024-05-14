package main

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/jmoiron/sqlx"
	_ "modernc.org/sqlite"
)

type TagVideo struct {
	Tags string
}

func getList(ctx context.Context, db *sqlx.DB) (videos []TagVideo, err error) {
	sqlSelect := `SELECT tags FROM videos 
                    WHERE tags LIKE '%worst%' GROUP BY tags`
	err = db.SelectContext(ctx, &videos, sqlSelect)
	return
}

func main() {
	db := sqlx.MustOpen("sqlite", "newvideo.db")
	defer db.Close()

	ctx := context.Background()
	videos, err := getList(ctx, db)
	if err != nil {
		log.Fatal(err)
	}

	var updates int64
	sqlUpdate := "UPDATE videos SET tags = ? WHERE tags = ?"

	for _, v := range videos {
		var tags []string
		// удаляем лишние теги
		for _, tag := range strings.Split(v.Tags, `|`) {
			if !strings.Contains(strings.ToLower(tag), "worst") {
				tags = append(tags, tag)
			}
		}

		res := db.MustExecContext(ctx, sqlUpdate, strings.Join(tags, `|`), v.Tags)
		if upd, err := res.RowsAffected(); err == nil {
			updates += upd
		}
	}
	fmt.Println(updates)
}
