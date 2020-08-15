package main

import (
	"net/http"
	"os"

	"entegrasyon/trendyol"

	"github.com/getsentry/sentry-go"
)

func main() {
	router := http.NewServeMux()
	staticFileServer := http.FileServer(http.Dir("templates/"))
	StaticFolderPath := "/templates/"

	router.Handle(StaticFolderPath, http.StripPrefix(StaticFolderPath, staticFileServer))
	router.HandleFunc("/", trendyol.GetLayout)
	if err := http.ListenAndServe(":8080", router); err != nil {
		sentry.CaptureException(err)
		os.Exit(1)
	}
}
