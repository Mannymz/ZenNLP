package go_sdk

import (
	"context"
	"fmt"
	"time"

	pb "github.com/Mannymz/ZenNLP/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

// Client provides a simplified interface to the NLP service
type Client struct {
	conn   *grpc.ClientConn
	client pb.NLPManagerClient
}

// Config holds client configuration options
type Config struct {
	Address    string
	Timeout    time.Duration
	MaxRetries int
}

// NewClient creates a new NLP client with the given address
func NewClient(addr string) (*Client, error) {
	return NewClientWithConfig(Config{
		Address:    addr,
		Timeout:    30 * time.Second,
		MaxRetries: 3,
	})
}

// NewClientWithConfig creates a new NLP client with custom configuration
func NewClientWithConfig(cfg Config) (*Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), cfg.Timeout)
	defer cancel()

	conn, err := grpc.DialContext(ctx, cfg.Address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to %s: %w", cfg.Address, err)
	}

	client := pb.NewNLPManagerClient(conn)
	return &Client{
		conn:   conn,
		client: client,
	}, nil
}

// Analyze performs sentiment analysis on the given text
func (c *Client) Analyze(ctx context.Context, text string) (*Result, error) {
	return c.AnalyzeWithLanguage(ctx, text, "fa")
}

// AnalyzeWithLanguage performs sentiment analysis on the given text with specified language
func (c *Client) AnalyzeWithLanguage(ctx context.Context, text, lang string) (*Result, error) {
	req := &pb.SentimentRequest{
		Text: text,
		Lang: lang,
	}

	resp, err := c.client.AnalyzeSentiment(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("sentiment analysis failed: %w", err)
	}

	return &Result{
		Label: resp.Label,
		Score: resp.Score,
	}, nil
}

// AnalyzeWithRetry performs sentiment analysis with automatic retries
func (c *Client) AnalyzeWithRetry(ctx context.Context, text string, maxRetries int) (*Result, error) {
	return c.AnalyzeWithLanguageAndRetry(ctx, text, "fa", maxRetries)
}

// AnalyzeWithLanguageAndRetry performs sentiment analysis with language and automatic retries
func (c *Client) AnalyzeWithLanguageAndRetry(ctx context.Context, text, lang string, maxRetries int) (*Result, error) {
	var lastErr error

	for attempt := 0; attempt <= maxRetries; attempt++ {
		if attempt > 0 {
			// Exponential backoff
			backoff := time.Duration(1<<uint(attempt-1)) * time.Second
			select {
			case <-ctx.Done():
				return nil, ctx.Err()
			case <-time.After(backoff):
			}
		}

		result, err := c.AnalyzeWithLanguage(ctx, text, lang)
		if err == nil {
			return result, nil
		}

		lastErr = err

		// Don't retry on certain error types
		if status, ok := status.FromError(err); ok {
			switch status.Code() {
			case codes.InvalidArgument, codes.PermissionDenied, codes.Unauthenticated:
				return nil, err
			}
		}
	}

	return nil, fmt.Errorf("max retries (%d) exceeded: %w", maxRetries, lastErr)
}

// Result represents the sentiment analysis result
type Result struct {
	Label string
	Score float64
}

// IsPositive returns true if the sentiment is positive
func (r *Result) IsPositive() bool {
	return r.Label == "positive"
}

// IsNegative returns true if the sentiment is negative
func (r *Result) IsNegative() bool {
	return r.Label == "negative"
}

// Confidence returns the confidence score as a percentage
func (r *Result) Confidence() float64 {
	return r.Score * 100
}

// Close closes the client connection
func (c *Client) Close() error {
	return c.conn.Close()
}
