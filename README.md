# Snack

Slackbot Proxy

### 環境構築

```
$ brew install protobuf
$ go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
$ protoc --go_out=api/grpc api/proto/hello.proto 
```
