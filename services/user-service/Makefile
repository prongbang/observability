build:
	cargo build

run_server:
	cargo run --bin user-server

run_client:
	cargo run --bin user-client

# brew install grpcurl
get_user:
	grpcurl -plaintext -import-path ./proto -proto user.proto -d '{"username": "em"}' '[::1]:50051' user.User/GetUser

# Response:
# 	{
#   	"id": "1",
#   	"name": "dev day",
#   	"username": "em",
#   	"password": "1234"
# 	}

build_image:
	docker build -t user-service .

run_image:
	docker run -p 50051:50051 user-service