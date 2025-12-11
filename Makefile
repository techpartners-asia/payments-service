

generate-proto:
	protoc --go_out=./pkg/proto/payment --go-grpc_out=./pkg/proto/payment --proto_path=./pkg/proto ./pkg/proto/payment.proto
	protoc --go_out=./pkg/proto/merchant --go-grpc_out=./pkg/proto/merchant --proto_path=./pkg/proto ./pkg/proto/merchant.proto