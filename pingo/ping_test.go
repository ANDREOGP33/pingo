package pingo_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ANDREOGP33/pingo/pingo"
)

func TestNormalizePort(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		port      string
		wantPort  string
		wantError bool
	}{
		{
			name:      "porta vazia",
			port:      "",
			wantError: false,
			wantPort: ":9876",
		},
		{
			name:     "porta sem dois pontos",
			port:     "8080",
			wantPort: ":8080",
		},
		{
			name:     "porta já com dois pontos",
			port:     ":8080",
			wantPort: ":8080",
		},
		{
			name:     "porta com espaços",
			port:     " 80 80 ",
			wantPort: ":8080",
		},
		{
			name:"porta com dois pontos no fim",
			port: "8080:",
			wantPort: ":8080",
		},
		{
			name:"porta com letras",
			port: "808a0",
			wantPort: ":8080",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := pingo.NormalizePort(tt.port)

			if tt.wantError {
				if err == nil {
					t.Fatalf("esperava erro, mas não ocorreu; got=%q", got)
				}
				return
			}

			if err != nil {
				t.Fatalf("não esperava erro, mas ocorreu: %v", err)
			}

			if got != tt.wantPort {
				t.Fatalf("NormalizePort(%q) = %q; queríamos %q", tt.port, got, tt.wantPort)
			}
		})
	}
}

func TestPingEndpoint(t *testing.T) {
	t.Parallel()

	mux := pingo.NewMux()

	req := httptest.NewRequest(http.MethodGet, "/ping", nil)
	rr := httptest.NewRecorder()

	mux.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("status code = %d; queríamos %d", rr.Code, http.StatusOK)
	}

	body := rr.Body.String()
	if body != "pong" {
		t.Fatalf("corpo da resposta = %q; queríamos %q", body, "pong")
	}
}
