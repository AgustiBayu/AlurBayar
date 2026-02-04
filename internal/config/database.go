package config

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

func NewDB() *sql.DB {
	user := GetEnv("DB_USER", "postgres")
	pass := GetEnv("DB_PASSWORD", "")
	host := GetEnv("DB_HOST", "localhost")
	port := GetEnv("DB_PORT", "5432")
	name := GetEnv("DB_NAME", "db_alurbayartest")
	ssl := GetEnv("DB_SSLMODE", "disable")

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", user, pass, host, port, name, ssl)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("failed conection in database: %v", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("database unreachable: %v", err)
	}

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db

}
