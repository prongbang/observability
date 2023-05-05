# brew install protobuf
# go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
# go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
# make gen name=auth
gen:
	protoc --go_out=proto --go_opt=paths=import \
        --go-grpc_out=proto --go-grpc_opt=paths=import \
        proto/$(name).proto

gen_auth:
	make gen name=auth

run_server:
	go run .

# brew install grpcurl
login:
	grpcurl -plaintext -import-path ./proto -proto auth.proto -d '{"username": "em", "password": "1234"}' '[::1]:50052' auth.Auth/Login

build_image:
	docker build -t auth-service .

run_image:
	docker run -p 50052:50052 auth-service