package main

import (
	"log"
	"net/http"
	"os"
)

func Log(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}

func main() {
	var port string

	if len(os.Args) <= 2 {
		port = "8000"
	} else {
		port = os.Args[1]
	}

	log.Printf("Listening on http://127.0.0.1:%s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, Log(http.FileServer(http.Dir(os.Getenv("PWD"))))))
}
