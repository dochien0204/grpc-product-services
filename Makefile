server:
	go run cmd/main.go
proto:
	protoc ./pkg/pb/*.proto --go_out=plugins=grpc:.