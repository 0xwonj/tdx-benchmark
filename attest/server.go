package attest

import (
	"context"
	"log"
	"time"

	pb "github.com/0xwonj/tdx-benchmark/proto/attest"
	"github.com/google/go-tdx-guest/client"
	tdxpb "github.com/google/go-tdx-guest/proto/tdx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Server implements the AttestService gRPC server
type Server struct {
	pb.UnimplementedAttestServiceServer
	quoteProvider client.QuoteProvider
	useMock       bool
}

// MockQuoteProvider implements a minimal mock of the QuoteProvider interface
// for testing purposes when TDX is not available
type MockQuoteProvider struct{}

// IsSupported always returns nil for the mock implementation
func (m *MockQuoteProvider) IsSupported() error {
	return nil
}

// GetRawQuote is required by the interface but not actually used in our mock
// implementation since we mock at the GetQuote level
func (m *MockQuoteProvider) GetRawQuote(reportData [64]byte) ([]byte, error) {
	return []byte("MOCK_TDX_QUOTE_DATA"), nil
}

// NewServer creates a new instance of the AttestService server
func NewServer() (*Server, error) {
	// Try to get the real quote provider from the go-tdx-guest library
	quoteProvider, err := client.GetQuoteProvider()
	if err != nil {
		log.Printf("Warning: Failed to get quote provider: %v. Using mock implementation.", err)
		return &Server{
			quoteProvider: &MockQuoteProvider{},
			useMock:       true,
		}, nil
	}

	// Check if TDX is supported
	if err := quoteProvider.IsSupported(); err != nil {
		log.Printf("Warning: TDX is not supported on this system: %v. Using mock implementation.", err)
		return &Server{
			quoteProvider: &MockQuoteProvider{},
			useMock:       true,
		}, nil
	}

	return &Server{
		quoteProvider: quoteProvider,
		useMock:       false,
	}, nil
}

// GetQuote implements the GetQuote RPC method
func (s *Server) GetQuote(ctx context.Context, req *pb.GetQuoteRequest) (*pb.GetQuoteResponse, error) {
	// The go-tdx-guest library expects a 64-byte report data
	var reportData [64]byte
	
	// Copy the report data into the fixed-size array
	// If report data is smaller than 64 bytes, the rest will be zeros
	// If it's larger, it will be truncated
	copy(reportData[:], req.ReportData)
	
	var quoteObj any
	var err error
	
	// Get the quote either from real provider or mock
	if s.useMock {
		// In a mock scenario, create a minimal QuoteV4 structure
		// Just enough to pass conversion to pb.Quote
		quoteObj = &tdxpb.QuoteV4{
			Header: &tdxpb.Header{
				Version:            4,
				AttestationKeyType: 2,
				TeeType:            0x00000081,
			},
			TdQuoteBody: &tdxpb.TDQuoteBody{
				ReportData: req.ReportData,
			},
			SignedDataSize: 64,
			SignedData: &tdxpb.Ecdsa256BitQuoteV4AuthData{
				Signature: make([]byte, 64),
			},
		}
		
		// Simulate some processing time to make the benchmark more realistic
		time.Sleep(5 * time.Millisecond)
	} else {
		// Get the real structured quote
		quoteObj, err = client.GetQuote(s.quoteProvider, reportData)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "failed to get quote: %v", err)
		}
	}

	// Cast the returned interface{} to QuoteV4
	quoteV4, ok := quoteObj.(*tdxpb.QuoteV4)
	if !ok {
		return nil, status.Errorf(codes.Internal, "unexpected quote format, expected QuoteV4")
	}
	
	// Convert the go-tdx-guest QuoteV4 to our proto Quote format using the utility function
	quote := ConvertQuoteV4ToQuote(quoteV4)
	
	return &pb.GetQuoteResponse{
		Quote: quote,
	}, nil
}
