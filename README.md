# Snack

<img src="./logo.png" width="500">

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
$ go install entgo.io/ent/cmd/ent@latest
$ ent new --target pkg/db/schema Webhook
$ ent generate ./pkg/db/schema --target pkg/db/ent
```
