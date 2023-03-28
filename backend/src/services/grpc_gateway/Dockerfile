# Step 1: build a go binary for the required service
ARG GO_VERSION=1.20.2
FROM golang:${GO_VERSION}-alpine AS builder

WORKDIR /app

COPY go.work go.work.sum ./

# TODO add all go.mod & go.sum in a betкter way
COPY ./src/services/grpc_gateway/go.mod ./src/services/grpc_gateway/go.sum ./src/services/grpc_gateway/
COPY ./src/services/users/go.mod ./src/services/users/go.sum ./src/services/users/
COPY ./src/services/houses/go.mod ./src/services/houses/go.sum ./src/services/houses/
COPY ./src/common/go.mod ./src/common/go.sum ./src/common/

RUN go mod download

COPY . .

RUN --mount=type=cache,target=/root/.cache/go-build CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags '-w -s' -installsuffix cgo tuhla/services/grpc_gateway

# Step 2: copy the built binary over to a new bare bone image
FROM scratch

COPY --from=builder /app/grpc_gateway /grpc_gateway

EXPOSE 1122 1123

ENTRYPOINT ["/grpc_gateway"]
