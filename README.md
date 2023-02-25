# Snack

<img src="./logo.png" width="100">

Slackbot Proxy

### 環境構築

```bash
$ brew install protobuf
$ go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
$ protoc --go_out=. --go-grpc_out=. api/health_check.proto
$ brew install grpcurl
$ grpcurl -plaintext localhost:8080 list
$ grpcurl -plaintext localhost:8080 list snack.HealthCheckService
$ grpcurl -plaintext -d '{"name": "snack"}' localhost:8080 snack.HealthCheckService.Echo
$ go install github.com/cosmtrek/air@latest
$ air
$ touch database.sqlite3
$ go install -tags 'sqlite3' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
$ migrate create -ext sql -dir pkg/db/migrations -seq create_webhooks_table
$ migrate -database sqlite3://database.sqlite3 -path pkg/db/migrations up
```
