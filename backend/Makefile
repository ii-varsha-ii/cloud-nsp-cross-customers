generate-protos:
	@echo "Generating protos..."
	protoc --go_out=./go_pb_files/ --go_opt=module=github.com/ii-varsha-ii/cloud-nsp-cross-customers/backend/protos --go-grpc_out=./go_pb_files/ --go-grpc_opt=module=github.com/ii-varsha-ii/cloud-nsp-cross-customers/backend/protos ./protos/organization.proto
	protoc --go_out=./go_pb_files/ --go_opt=module=github.com/ii-varsha-ii/cloud-nsp-cross-customers/backend/protos --go-grpc_out=./go_pb_files/ --go-grpc_opt=module=github.com/ii-varsha-ii/cloud-nsp-cross-customers/backend/protos ./protos/vpc.proto
	protoc --go_out=./go_pb_files/ --go_opt=module=github.com/ii-varsha-ii/cloud-nsp-cross-customers/backend/protos --go-grpc_out=./go_pb_files/ --go-grpc_opt=module=github.com/ii-varsha-ii/cloud-nsp-cross-customers/backend/protos ./protos/gateway.proto
	@echo "Generation of protos complete"