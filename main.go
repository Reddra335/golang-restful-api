package main

import (
	"net/http"
	"restful_api/middleware"

	_ "github.com/go-sql-driver/mysql"
)

func NewServer(authMiddleWare *middleware.AuthMidleware) *http.Server {

	return &http.Server{
		Addr:    "localhost:8080",
		Handler: authMiddleWare,
	}
}

func main() {

	// Server := NewServer(authMiddleWare)
	// err := server.ListenAndServe()
	// helper.ErrorT(err)

}
