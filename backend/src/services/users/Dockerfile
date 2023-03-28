ARG GO_VERSION=1.20.2
FROM golang:${GO_VERSION}-alpine AS builder

WORKDIR /app

COPY go.work go.work.sum ./

# TODO add all go.mod & go.sum in a better way
COPY ./src/services/users/go.mod ./src/services/users/go.sum ./src/services/users/
COPY ./src/services/houses/go.mod ./src/services/houses/go.sum ./src/services/houses/
COPY ./src/services/grpc_gateway/go.mod ./src/services/grpc_gateway/go.sum ./src/services/grpc_gateway/
COPY ./src/common/go.mod ./src/common/go.sum ./src/common/

RUN go mod download

COPY . .

RUN --mount=type=cache,target=/root/.cache/go-build CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags '-w -s' -installsuffix cgo tuhla/services/users

# Copy the binary to a new image.
FROM scratch

COPY --from=builder /app/users /users

EXPOSE 1124

ENTRYPOINT ["/users"]
