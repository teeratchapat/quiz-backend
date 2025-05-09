package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func Connect(url string) *sql.DB {
	db, err := sql.Open("postgres", url)
	if err != nil {
		log.Println("SQL Open error:", err)
		return nil
	}
	if err = db.Ping(); err != nil {
		log.Println("❌ DB connection failed:", err)
		return nil
	}

	log.Println("✅ DB connected successfully")
	return db
}
