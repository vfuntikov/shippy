gen:
	protoc --proto_path=./proto/vessel --micro_out=. --go_out=. proto/vessel/vessel.proto
build:
	docker build -t vessel-service .
run:
	docker run -p 50052:50051 -e MICRO_SERVER_ADDRESS=:50051 -e MICRO_REGISTRY=mdns vessel-service