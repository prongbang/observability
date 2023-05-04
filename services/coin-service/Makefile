
# python3 -m pip install grpcio-tools
gen:
	python3 -m grpc_tools.protoc -I./proto --python_out=. --grpc_python_out=. ./proto/coin.proto

run_server:
	python3 main.py

# brew install grpcurl
get_coin:
	grpcurl -plaintext -import-path ./proto -proto coin.proto -d '{"username": "em"}' '[::1]:50053' coin.Coin/GetCoin

build_image:
	docker build -t coin-service .

run_image:
	docker run -p 50053:50053 coin-service