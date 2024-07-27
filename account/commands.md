```bash=
gzip -dk *.gz
protoc --proto_path=proto --go_out=. --go_opt=module=github.com/Nu1LSpaxe/grpc-go-for-professionals/account proto/*.proto
go run main.go
```

```bash=
protoc --proto_path=proto \
  --go_out=. --go_opt=module=github.com/Nu1LSpaxe/grpc-go-for-professionals/account \
  --go-grpc_out=. --go-grpc_opt=module=github.com/Nu1LSpaxe/grpc-go-for-professionals/account \
  proto/*.proto
```