package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Mannymz/ZenNLP/go-sdk"
)

func main() {
	// Create a new client
	client, err := go_sdk.NewClient("localhost:50051")
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	defer client.Close()

	// Persian text examples
	persianTexts := []string{
		"این محصول عالی است و من خیلی راضی هستم",
		"کیفیت بسیار بد بود و اصلا توصیه نمی‌کنم",
		"قیمت مناسبی دارد و کیفیت خوبی هم دارد",
		"بسیار ناامید شدم، این محصول ارزش خرید ندارد",
	}

	fmt.Println("=== ZenNLP Persian Sentiment Analysis ===")
	fmt.Println()

	// Analyze each text with context timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	for i, text := range persianTexts {
		fmt.Printf("Text %d: %s\n", i+1, text)

		// Simple analysis
		result, err := client.Analyze(ctx, text)
		if err != nil {
			log.Printf("Error analyzing text %d: %v", i+1, err)
			continue
		}

		fmt.Printf("  Sentiment: %s\n", result.Label)
		fmt.Printf("  Confidence: %.2f%%\n", result.Confidence())
		fmt.Printf("  Is Positive: %t\n", result.IsPositive())
		fmt.Printf("  Is Negative: %t\n", result.IsNegative())
		fmt.Println()

		// Demonstrate retry functionality
		retryResult, err := client.AnalyzeWithRetry(ctx, text, 2)
		if err != nil {
			log.Printf("Retry failed for text %d: %v", i+1, err)
		} else {
			fmt.Printf("  Retry Result: %s (%.2f%%)\n", retryResult.Label, retryResult.Confidence())
		}
		fmt.Println("---")
	}

	// Demonstrate custom configuration
	fmt.Println("=== Custom Configuration Example ===")
	customClient, err := go_sdk.NewClientWithConfig(go_sdk.Config{
		Address:    "localhost:50051",
		Timeout:    5 * time.Second,
		MaxRetries: 1,
	})
	if err != nil {
		log.Printf("Failed to create custom client: %v", err)
		return
	}
	defer customClient.Close()

	testText := "تجربه کار با این سرویس بسیار خوب بود"
	result, err := customClient.Analyze(ctx, testText)
	if err != nil {
		log.Printf("Custom client error: %v", err)
		return
	}

	fmt.Printf("Custom Client Result: %s (%.2f%% confidence)\n", result.Label, result.Confidence())
}