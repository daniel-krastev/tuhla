#!/bin/bash

# Generates protobuf and gRPC required files within the current directory using protoc command.

# This is 
PREFIX="tuhla/services/users/proto"
COMMAND="protoc --proto_path=. --go_out=. --go_opt=module=${PREFIX} --go-grpc_out=. --go-grpc_opt=module=${PREFIX} *.proto"

PROJECT_ROOT="/home/dani/work/tuhla"
SERVICE_PROTO_FILE="/backend/src/services/users/proto"
CURRENT_DIR=$PWD

# Execute the command within the proto folder.
cd "${PROJECT_ROOT}${SERVICE_PROTO_FILE}" && eval $COMMAND

# Return to the current folder.
cd $CURRENT_DIR
