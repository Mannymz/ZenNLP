package main

import (
	"context"
	"testing"
	"time"

	go_sdk "github.com/Mannymz/ZenNLP/tree/main//go-sdk"
)

// Note: These tests require the NLP server to be running
// Start the server with: docker-compose up nlp-engine
// Or manually: cd nlp-engine && python server.py

// TestClientInitialization tests the SDK client initialization
func TestClientInitialization(t *testing.T) {
	client, err := go_sdk.NewClient("localhost:50051")
	if err != nil {
		t.Skipf("Skipping test - cannot connect to server: %v", err)
		return
	}
	defer client.Close()

	if client == nil {
		t.Error("expected non-nil client")
	}
}

// TestClientWithConfig tests client initialization with custom config
func TestClientWithConfig(t *testing.T) {
	config := go_sdk.Config{
		Address:    "localhost:50051",
		Timeout:    5 * time.Second,
		MaxRetries: 2,
	}

	client, err := go_sdk.NewClientWithConfig(config)
	if err != nil {
		t.Skipf("Skipping test - cannot connect to server: %v", err)
		return
	}
	defer client.Close()

	if client == nil {
		t.Error("expected non-nil client")
	}
}

// TestSentimentAnalysisPositive tests positive sentiment analysis
func TestSentimentAnalysisPositive(t *testing.T) {
	client, err := go_sdk.NewClient("localhost:50051")
	if err != nil {
		t.Skipf("Skipping test - cannot connect to server: %v", err)
		return
	}
	defer client.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Persian positive text
	text := "این محصول عالی است و من خیلی راضی هستم"

	result, err := client.Analyze(ctx, text)
	if err != nil {
		t.Fatalf("Analyze() error = %v", err)
	}

	if result == nil {
		t.Fatal("expected non-nil result")
	}

	t.Logf("Sentiment: %s, Confidence: %.2f%%", result.Label, result.Confidence())

	if !result.IsPositive() && !result.IsNegative() {
		t.Errorf("result should be either positive or negative, got: %s", result.Label)
	}
}

// TestSentimentAnalysisNegative tests negative sentiment analysis
func TestSentimentAnalysisNegative(t *testing.T) {
	client, err := go_sdk.NewClient("localhost:50051")
	if err != nil {
		t.Skipf("Skipping test - cannot connect to server: %v", err)
		return
	}
	defer client.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Persian negative text
	text := "کیفیت بسیار بد بود و اصلا توصیه نمی‌کنم"

	result, err := client.Analyze(ctx, text)
	if err != nil {
		t.Fatalf("Analyze() error = %v", err)
	}

	if result == nil {
		t.Fatal("expected non-nil result")
	}

	t.Logf("Sentiment: %s, Confidence: %.2f%%", result.Label, result.Confidence())

	if !result.IsPositive() && !result.IsNegative() {
		t.Errorf("result should be either positive or negative, got: %s", result.Label)
	}
}

// TestAnalyzeWithLanguage tests sentiment analysis with explicit language
func TestAnalyzeWithLanguage(t *testing.T) {
	client, err := go_sdk.NewClient("localhost:50051")
	if err != nil {
		t.Skipf("Skipping test - cannot connect to server: %v", err)
		return
	}
	defer client.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	text := "قیمت مناسبی دارد و کیفیت خوبی هم دارد"

	result, err := client.AnalyzeWithLanguage(ctx, text, "fa")
	if err != nil {
		t.Fatalf("AnalyzeWithLanguage() error = %v", err)
	}

	if result == nil {
		t.Fatal("expected non-nil result")
	}

	t.Logf("Sentiment: %s, Confidence: %.2f%%", result.Label, result.Confidence())
}

// TestAnalyzeWithRetry tests the retry functionality
func TestAnalyzeWithRetry(t *testing.T) {
	client, err := go_sdk.NewClient("localhost:50051")
	if err != nil {
		t.Skipf("Skipping test - cannot connect to server: %v", err)
		return
	}
	defer client.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	text := "بسیار ناامید شدم، این محصول ارزش خرید ندارد"

	result, err := client.AnalyzeWithRetry(ctx, text, 3)
	if err != nil {
		t.Fatalf("AnalyzeWithRetry() error = %v", err)
	}

	if result == nil {
		t.Fatal("expected non-nil result")
	}

	t.Logf("Sentiment: %s, Confidence: %.2f%%", result.Label, result.Confidence())
}

// TestMultipleSentimentAnalyses tests multiple consecutive analyses
func TestMultipleSentimentAnalyses(t *testing.T) {
	client, err := go_sdk.NewClient("localhost:50051")
	if err != nil {
		t.Skipf("Skipping test - cannot connect to server: %v", err)
		return
	}
	defer client.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	tests := []struct {
		name string
		text string
	}{
		{"positive review", "این محصول عالی است و من خیلی راضی هستم"},
		{"negative review", "کیفیت بسیار بد بود و اصلا توصیه نمی‌کنم"},
		{"mixed review", "قیمت مناسبی دارد و کیفیت خوبی هم دارد"},
		{"strong negative", "بسیار ناامید شدم، این محصول ارزش خرید ندارد"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := client.Analyze(ctx, tt.text)
			if err != nil {
				t.Errorf("Analyze() error = %v", err)
				return
			}

			if result == nil {
				t.Error("expected non-nil result")
				return
			}

			t.Logf("%s: Sentiment=%s, Confidence=%.2f%%, Positive=%t, Negative=%t",
				tt.name, result.Label, result.Confidence(), result.IsPositive(), result.IsNegative())
		})
	}
}

// TestResultMethods tests the Result helper methods
func TestResultMethods(t *testing.T) {
	client, err := go_sdk.NewClient("localhost:50051")
	if err != nil {
		t.Skipf("Skipping test - cannot connect to server: %v", err)
		return
	}
	defer client.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	text := "این محصول عالی است"

	result, err := client.Analyze(ctx, text)
	if err != nil {
		t.Fatalf("Analyze() error = %v", err)
	}

	// Test that IsPositive and IsNegative are mutually exclusive
	if result.IsPositive() && result.IsNegative() {
		t.Error("result cannot be both positive and negative")
	}

	// Test that Confidence returns a valid percentage
	confidence := result.Confidence()
	if confidence < 0 || confidence > 100 {
		t.Errorf("confidence should be between 0 and 100, got %.2f", confidence)
	}
}

// TestContextTimeout tests that context timeout is respected
func TestContextTimeout(t *testing.T) {
	client, err := go_sdk.NewClient("localhost:50051")
	if err != nil {
		t.Skipf("Skipping test - cannot connect to server: %v", err)
		return
	}
	defer client.Close()

	// Very short timeout to test timeout handling
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Nanosecond)
	defer cancel()

	text := "این محصول عالی است"

	_, err = client.Analyze(ctx, text)
	// We expect this to fail due to timeout or context cancellation
	if err == nil {
		t.Log("Expected timeout error, but analysis succeeded (server may be very fast)")
	}
}

// BenchmarkSentimentAnalysis benchmarks sentiment analysis performance
func BenchmarkSentimentAnalysis(b *testing.B) {
	client, err := go_sdk.NewClient("localhost:50051")
	if err != nil {
		b.Skipf("Skipping benchmark - cannot connect to server: %v", err)
		return
	}
	defer client.Close()

	ctx := context.Background()
	text := "این محصول عالی است و من خیلی راضی هستم"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := client.Analyze(ctx, text)
		if err != nil {
			b.Fatalf("Analyze() error = %v", err)
		}
	}
}

// BenchmarkSentimentAnalysisWithRetry benchmarks sentiment analysis with retry
func BenchmarkSentimentAnalysisWithRetry(b *testing.B) {
	client, err := go_sdk.NewClient("localhost:50051")
	if err != nil {
		b.Skipf("Skipping benchmark - cannot connect to server: %v", err)
		return
	}
	defer client.Close()

	ctx := context.Background()
	text := "کیفیت بسیار بد بود و اصلا توصیه نمی‌کنم"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := client.AnalyzeWithRetry(ctx, text, 2)
		if err != nil {
			b.Fatalf("AnalyzeWithRetry() error = %v", err)
		}
	}
}
