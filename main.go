package main

import (
	log "github.com/sirupsen/logrus"
	"net/http"

	"entegrasyon/trendyol"
)

func main() {
	router := http.NewServeMux()
	staticFileServer := http.FileServer(http.Dir("templates/"))
	StaticFolderPath := "/templates/"

	router.Handle(StaticFolderPath, http.StripPrefix(StaticFolderPath, staticFileServer))
	router.HandleFunc("/", trendyol.GetLayout)
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Error(err)
	}
}
