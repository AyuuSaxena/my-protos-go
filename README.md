# proto compiler
make sure the 'protoc' compiler is installed
https://grpc.io/docs/protoc-installation/

# Install following dependencies for go proto compiler
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

# For protofile linting 
go install github.com/yoheimuta/protolint/cmd/protolint@latest

# Install following dependencies for js code
npm install grpc_tools_node_protoc_ts 
npm install grpc-tools
npm install ts-proto

# Run linting
protolint lint .

# Generate the code
/bin/bash generate.sh

