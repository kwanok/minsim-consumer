.PHONY: generate-minsim-proto
generate-minsim-proto:
	protoc -I . \
		--go_out . --go_opt paths=source_relative \
		--go-grpc_out . --go-grpc_opt paths=source_relative \
		service/minsim/minsim.proto


.PHONY: server
server:
	go run main.go --config=.config.yaml