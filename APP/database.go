package app

import (
	"database/sql"
	"fmt"
	"os"
	"restful_api/helper"
	"time"
)

func NewDB() *sql.DB {
	host := getEnv("DB_HOST", "localhost")
	port := getEnv("DB_PORT", "3306")
	user := getEnv("DB_USER", "root")
	pass := getEnv("DB_PASSWORD", "")
	name := getEnv("DB_NAME", "belejar_golang_restful_api")

	auth := user
	if pass != "" {
		auth = user + ":" + pass
	}

	dsn := fmt.Sprintf("%s@tcp(%s:%s)/%s?parseTime=true", auth, host, port, name)

	db, err := sql.Open("mysql", dsn)
	helper.ErrorT(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Hour)
	db.SetConnMaxIdleTime(10 * time.Hour)

	return db
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
