package pingo

import (
	"log"
	"net/http"
)

func InitServer(port string) {
	mux := http.NewServeMux()

	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("pong"))
	})

	if err := http.ListenAndServe(port, mux); err != nil {
		log.Fatal(err)
	}
}
