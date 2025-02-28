package attest

import (
	"context"
	"fmt"

	pb "github.com/0xwonj/tdx-benchmark/proto/attest"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// Client is a client for the AttestService
type Client struct {
	client pb.AttestServiceClient
	conn   *grpc.ClientConn
}

// NewClient creates a new client that connects to the AttestService at the specified address
func NewClient(serverAddr string) (*Client, error) {
	// Set up connection options
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	// Establish the connection
	conn, err := grpc.NewClient(serverAddr, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to server: %v", err)
	}
	
	return &Client{
		client: pb.NewAttestServiceClient(conn),
		conn:   conn,
	}, nil
}

// Close closes the client connection
func (c *Client) Close() error {
	return c.conn.Close()
}

// GetQuote requests a TDX quote from the server with the specified report data
func (c *Client) GetQuote(ctx context.Context, reportData []byte) (*pb.Quote, error) {
	// Create a GetQuoteRequest with the provided report data
	req := &pb.GetQuoteRequest{
		ReportData: reportData,
	}
	
	// Call the server's GetQuote method
	resp, err := c.client.GetQuote(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to get quote: %v", err)
	}
	
	return resp.Quote, nil
}
