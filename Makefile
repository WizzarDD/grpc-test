.DEFAULT_GOAL := gen

gen:
	protoc --go_out=. --go_opt=paths=import \
	--go-grpc_out=. --go-grpc_opt=paths=import proto/balances.proto