#!/bin/bash

# Generates protobuf and gRPC required files within the current directory using protoc command.

# Removes this part of the go_package, so the go files are not nested in this structure.
PREFIX="tuhla/services/grpc_gateway/proto"

PROTOC_COMMAND="\
protoc \
--proto_path=. --proto_path=$GOPATH/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.16.0/third_party/googleapis \
--go_out=. --go_opt=module=${PREFIX} \
--go-grpc_out=. --go-grpc_opt=module=${PREFIX} \
--grpc-gateway_out=. --grpc-gateway_opt=module=${PREFIX} \
*/*.proto"

PROJECT_ROOT="/home/dani/work/tuhla"
SERVICE_PROTO_FILE="/backend/src/services/grpc_gateway/proto"
CURRENT_DIR=$PWD

# Execute the command within the proto folder.
cd "${PROJECT_ROOT}${SERVICE_PROTO_FILE}" && eval $PROTOC_COMMAND

# Return to the current folder.
cd $CURRENT_DIR
