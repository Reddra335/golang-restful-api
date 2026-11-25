package main

import (
	"net/http"
	"os"
	"restful_api/helper"
	"restful_api/middleware"

	_ "github.com/go-sql-driver/mysql"
)

func NewServer(authMiddleWare *middleware.AuthMidleware) *http.Server {
	addr := os.Getenv("APP_ADDR")
	if addr == "" {
		addr = ":8080"
	}

	return &http.Server{
		Addr:    addr,
		Handler: authMiddleWare,
	}
}

func main() {
	server := InitializedServer()
	err := server.ListenAndServe()
	helper.ErrorT(err)
}
