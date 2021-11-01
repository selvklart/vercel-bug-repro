package main

import (
	"net/http"

	"github.com/selvklart/vercel-bug-repro/api"
)

func main() {
	server := http.Server{
		Addr:    "localhost:3000",
		Handler: http.HandlerFunc(api.BugHandler),
	}

	server.ListenAndServe()
}
