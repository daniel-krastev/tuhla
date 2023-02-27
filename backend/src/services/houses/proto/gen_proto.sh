#!/bin/bash

# Generates protobuf and gRPC required files within the current directory using protoc command.

# Removes this part of the go_package, so the go files are not nested in this structure.
PREFIX="tuhla/services/houses/proto"
COMMAND="protoc --proto_path=. --go_out=. --go_opt=module=${PREFIX} --go-grpc_out=. --go-grpc_opt=module=${PREFIX} *.proto"

PROJECT_ROOT="/home/dani/work/tuhla"
SERVICE_PROTO_FILE="/backend/src/services/houses/proto"
CURRENT_DIR=$PWD

# Execute the command within the proto folder.
cd "${PROJECT_ROOT}${SERVICE_PROTO_FILE}" && eval $COMMAND

# Return to the current folder.
cd $CURRENT_DIR
