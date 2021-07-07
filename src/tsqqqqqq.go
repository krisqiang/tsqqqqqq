package main

import (
	"net/http"
	"time"
	"tsqqqqqq/router"
	"tsqqqqqq/util"
)

func main() {
	r := router.InitRouter()
	r.Use(util.Cors())
	s := &http.Server{
		Addr:           ":8080",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
