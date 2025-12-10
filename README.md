## pingo

Uma lib/mini-servidor HTTP pensada para ser embutida em vários serviços e projetos como verificador simples de status (healthcheck). A ideia é evitar entrar manualmente nos servidores só para checar se as aplicações estão de pé: basta consultar o endpoint de ping.


## Funcionalidades

- Endpoint leve de saúde: `GET /ping` → responde `200 OK` com corpo `pong`
- Funções utilitárias para integrar com servidores existentes:
  - `NormalizePort(port string) (string, error)`: normaliza/valida porta (ex.: `"8080"` → `":8080"`)
  - `NewMux() *http.ServeMux`: retorna um `ServeMux` já com o endpoint `/ping`
  - `InitServer(port string)`: inicia um servidor somente com o `/ping`


## Instalação

```bash
go get github.com/ANDREOGP33/pingo/pingo
```

Requer Go 1.20+ (recomendado).


## Uso rápido

### 1) Servidor standalone

```go
package main

import (
	"github.com/ANDREOGP33/pingo/pingo"
)

func main() {
	// Aceita "8080" ou ":8080"; a porta é normalizada internamente.
	pingo.InitServer("8080")
}
```

Teste rápido:

```bash
curl -i http://localhost:8080/ping
# HTTP/1.1 200 OK
# ...
# pong
```

### 2) Embutido em um servidor existente

```go
package main

import (
	"log"
	"net/http"

	"github.com/ANDREOGP33/pingo/pingo"
)

func main() {
	// Mux com /ping pronto
	healthMux := pingo.NewMux()

	// Integre no seu roteador/servidor atual
	http.Handle("/health/", http.StripPrefix("/health", healthMux))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
```

Agora o ping fica em `GET /health/ping`.


## Endpoints

- `GET /ping` → `200 OK` com corpo `pong`


## Testes

```bash
go test ./...
```

Os testes cobrem:
- Normalização/validação de porta (`NormalizePort`)
- Comportamento do endpoint `/ping` via `httptest`


## Integração com verificadores de status

Você pode apontar seu orquestrador/monitor (Kubernetes, Docker, Nginx, HAProxy, Consul, etc.) para `GET /ping`. É uma verificação rápida que indica se o processo está aceitando requisições.

Exemplos de uso:
- Liveness/Readiness Probe (K8s): `httpGet` em `/ping`
- Balanceadores: healthcheck HTTP `200` em `/ping`


## Roadmap (sugestões)

- Configurar prefixo customizável para o ping (ex.: `/health/ping`)
- Métricas e versão da aplicação em endpoint separado (ex.: `/health/info`)
- Opção de registrar handlers adicionais no `NewMux`


## Licença

A definir pelo autor do repositório.


