package main

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/0xwonj/tdx-benchmark/attest"
	"github.com/0xwonj/tdx-benchmark/benchmark"
	attestpb "github.com/0xwonj/tdx-benchmark/proto/attest"
	benchmarkpb "github.com/0xwonj/tdx-benchmark/proto/benchmark"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Listen on the specified port
	lis, err := net.Listen("tcp", ":"+os.Getenv("PORT"))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	log.Printf("Server listening on port %s in %s environment", lis.Addr(), os.Getenv("ENV"))

	// Create gRPC server and TDX client
	grpcServer := grpc.NewServer()

	// Create and register benchmark service
	benchmarkServer, err := benchmark.NewServer()
	if err != nil {
		log.Fatalf("Failed to create benchmark server: %v", err)
	}
	benchmarkpb.RegisterBenchmarkServiceServer(grpcServer, benchmarkServer)

	// Create and register attest service
	attestServer, err := attest.NewServer()
	if err != nil {
		log.Fatalf("Failed to create attest server: %v", err)
	}
	attestpb.RegisterAttestServiceServer(grpcServer, attestServer)

	// Enable reflection for debugging
	reflection.Register(grpcServer)

	// Graceful shutdown setup
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	// Run the server in a goroutine
	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()

	// Wait for interrupt signal
	<-stop
	log.Println("Shutting down server gracefully...")

	// Create a context with timeout for shutdown
	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()


	// Gracefully stop the gRPC server
	grpcServer.GracefulStop()

	// Perform any additional cleanup tasks if necessary
	log.Println("Server stopped")
}
