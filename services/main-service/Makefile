# go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
# go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
# make gen name=auth
gen:
	protoc --go_out=proto --go_opt=paths=import \
        --go-grpc_out=proto --go-grpc_opt=paths=import \
        proto/$(name).proto

gen_auth:
	make gen name=auth

gen_user:
	make gen name=user

run_server:
	go run .

build_image:
	docker build -t main-service .

run_image:
	docker run -p 3000:3000 main-service

get_login:
	curl --request GET -sL \
	     --url 'http://localhost:3000/v1/login'

get_user:
	curl --request GET -sL \
	     --url 'http://localhost:3000/v1/user'

get_coin:
	curl --request GET -sL \
	     --url 'http://localhost:3000/v1/coin'