package main

import (
	"database/sql"
	"net/http"
	"time"
	"tsqqqqqq/router"
)

var Db *sql.DB

func main() {
	r := router.InitRouter()
	s := &http.Server{
		Addr:           ":8080",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}
