package pingo

import (
	"log"
	"net/http"
	"regexp"
	"strconv"
)

var nonDigit = regexp.MustCompile(`\D`)

func NormalizePort(port string) (string, error) {
	if port == "" {
		return ":9876", nil
	}

	port = nonDigit.ReplaceAllString(port, "")

	n, err := strconv.Atoi(port)
	if n < 1 || n > 65535 {
		return ":9876", err
	}

	return ":" + port, nil
}

// NewMux devolve um mux HTTP configurado com o endpoint /ping.
func NewMux() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("pong"))
	})

	return mux
}

// InitServer inicia o servidor HTTP escutando na porta informada, caso a porta seja "" o servidor irá usar a porta 9876
func InitServer(port string) {
	normalizedPort, err := NormalizePort(port)
	if err != nil {
		log.Fatal(err)
	}

	mux := NewMux()

	if err := http.ListenAndServe(normalizedPort, mux); err != nil {
		log.Fatal(err)
	}
}
