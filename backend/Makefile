generate-protos:
	@echo "Generating protos..."
	protoc --go_out=./backend/go_pb_files/ --go_opt=module=github.com/ii-varsha-ii/nsp-cross-customer-multicloud/backend/protos --go-grpc_out=./backend/go_pb_files/ --go-grpc_opt=module=github.com/ii-varsha-ii/nsp-cross-customer-multicloud/backend/protos ./backend/protos/organization.proto
	@echo "Generation of protos complete"