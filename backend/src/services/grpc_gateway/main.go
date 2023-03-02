package main

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"strings"

	"tuhla/services/grpc_gateway/controller"
	"tuhla/services/grpc_gateway/proto/houses/housespb"
	"tuhla/services/grpc_gateway/proto/users/userspb"
	"tuhla/services/houses/proto/housesservicepb"
	"tuhla/services/users/proto/usersservicepb"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/reflection"
)

var (
	serviceAddress  = flag.String("service_address", "localhost", "service address")
	gRPCservicePort = flag.Int("grpc_service_port", 1123, "grpc service port")
	HTTPservicePort = flag.Int("http_service_port", 1122, "http service port")

	// users address
	usersAddress = flag.String("users_address", "localhost", "users address")
	usersPort    = flag.Int("users_port", 1124, "users port")

	// houses address
	housesAddress = flag.String("houses_address", "localhost", "houses address")
	housesPort    = flag.Int("houses_port", 1125, "houses port")

	log = grpclog.NewLoggerV2(os.Stdout, ioutil.Discard, ioutil.Discard)
)

func init() {
	grpclog.SetLoggerV2(log)
}

func main() {
	flag.Parse()

	// Create internal service clients.
	usersConn, err := grpc.Dial(fmt.Sprintf("%s:%d", *usersAddress, *usersPort), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("cannot connect to users service: %v", err)
	}
	usersClient := usersservicepb.NewUsersClient(usersConn)

	housesConn, err := grpc.Dial(fmt.Sprintf("%s:%d", *housesAddress, *housesPort), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("cannot connect to houses service: %v", err)
	}
	housesClient := housesservicepb.NewHousesClient(housesConn)

	controller := controller.New(usersClient, housesClient)

	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", *serviceAddress, *gRPCservicePort))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Cannot listen on the localhost: %v\n", err)
		os.Exit(1)
	}

	grpcServer := grpc.NewServer([]grpc.ServerOption{}...)
	reflection.Register(grpcServer)
	userspb.RegisterUsersServer(grpcServer, controller)
	housespb.RegisterHousesServer(grpcServer, controller)

	go func() {
		log.Fatal(grpcServer.Serve(lis))
	}()
	fmt.Println("gRPC service started...")

	log.Fatal(run())
}

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := userspb.RegisterUsersHandlerFromEndpoint(ctx, mux, fmt.Sprintf("%s:%d", *serviceAddress, *gRPCservicePort), opts)
	if err != nil {
		return err
	}
	err = housespb.RegisterHousesHandlerFromEndpoint(ctx, mux, fmt.Sprintf("%s:%d", *serviceAddress, *gRPCservicePort), opts)
	if err != nil {
		return err
	}

	httpServer := &http.Server{
		Addr: fmt.Sprintf("%s:%d", *serviceAddress, *HTTPservicePort),
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Dump the request body in a human-readable format
			requestDump, err := httputil.DumpRequest(r, true)
			if err != nil {
				http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
				return
			}
			// Print the request body to standard output
			fmt.Println(string(requestDump))

			// Serve CORS requests
			if origin := r.Header.Get("Origin"); origin != "" {
				w.Header().Set("Access-Control-Allow-Origin", origin)
				if r.Method == "OPTIONS" && r.Header.Get("Access-Control-Request-Method") != "" {
					preflightHandler(w, r)
					return
				}
			}

			mux.ServeHTTP(w, r)
		}),
	}

	return httpServer.ListenAndServe()
}

// preflightHandler adds the necessary headers in order to serve
// CORS from any origin using the methods "GET", "HEAD", "POST", "PUT", "DELETE"
// We insist, don't do this without consideration in production systems.
func preflightHandler(w http.ResponseWriter, r *http.Request) {
	headers := []string{"Content-Type", "Accept", "Authorization"}
	w.Header().Set("Access-Control-Allow-Headers", strings.Join(headers, ","))
	methods := []string{"GET", "HEAD", "POST", "PUT", "DELETE"}
	w.Header().Set("Access-Control-Allow-Methods", strings.Join(methods, ","))
	log.Infof("preflight request for %s", r.URL.Path)
}
