gen:
	# protoc -I. --go_out=plugins=grpc:. proto/consignment/consignment.proto
	protoc --proto_path=./proto/consignment --micro_out=. --go_out=. proto/consignment/consignment.proto
build:
	GOOS=linux GOARCH=amd64 go build -o app main.go
	docker build -t shippy-service-consignment .
run:
	docker run -p 50051:50051 -e MICRO_SERVER_ADDRESS=:50051 shippy-service-consignment
	