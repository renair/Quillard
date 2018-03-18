package main

import (
	"net/http"
	"strings"
)

func getFileHandler(prefix string, dir string) http.HandlerFunc {
	fileHandler := http.StripPrefix(prefix, http.FileServer(http.Dir(dir)))
	return func(resp http.ResponseWriter, req *http.Request) {
		url := req.URL.Path
		if url != "/" && strings.HasSuffix(url, "/") {
			http.NotFound(resp, req)
		} else {
			fileHandler.ServeHTTP(resp, req)
		}
	}
}
