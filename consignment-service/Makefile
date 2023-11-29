gen:
	protoc --proto_path=./proto/consignment --micro_out=. --go_out=. proto/consignment/consignment.proto
build:
	docker build -t consignment-service .
run:	
	docker run -p 50051:50051 -e MICRO_SERVER_ADDRESS=:50051 -e MICRO_REGISTRY=mdns consignment-service
mongo:
	docker run --name mongodb -e MONGO_INITDB_ROOT_USERNAME=admin -e MONGO_INITDB_ROOT_PASSWORD=mongo -p 27017:27017 mongo
