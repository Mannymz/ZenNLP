package examples

import (
	"fmt"
	"testing"

	"github.com/Mannymz/ZenNLP/sdk"
)

// TestInitializeClient tests the SDK client initialization
func TestInitializeClient(t *testing.T) {
	client, err := sdk.NewClient()
	if err != nil {
		t.Fatalf("failed to initialize client: %v", err)
	}

	if client == nil {
		t.Error("expected non-nil client")
	}
}

// TestTokenization tests the tokenization functionality
func TestTokenization(t *testing.T) {
	client, err := sdk.NewClient()
	if err != nil {
		t.Fatalf("failed to initialize client: %v", err)
	}

	tests := []struct {
		name    string
		input   string
		wantErr bool
	}{
		{
			name:    "simple sentence",
			input:   "Hello world",
			wantErr: false,
		},
		{
			name:    "sentence with punctuation",
			input:   "How are you?",
			wantErr: false,
		},
		{
			name:    "empty string",
			input:   "",
			wantErr: false,
		},
		{
			name:    "complex sentence",
			input:   "The quick brown fox jumps over the lazy dog.",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tokens, err := client.Tokenize(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Tokenize() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr && tokens == nil {
				t.Error("expected non-nil tokens")
			}
		})
	}
}

// TestSentenceSegmentation tests sentence segmentation
func TestSentenceSegmentation(t *testing.T) {
	client, err := sdk.NewClient()
	if err != nil {
		t.Fatalf("failed to initialize client: %v", err)
	}

	text := "This is the first sentence. This is the second sentence! And this is the third?"

	sentences, err := client.SegmentSentences(text)
	if err != nil {
		t.Fatalf("SegmentSentences() error = %v", err)
	}

	if sentences == nil {
		t.Error("expected non-nil sentences")
	}

	if len(sentences) == 0 {
		t.Error("expected at least one sentence")
	}
}

// TestPOSTagging tests Part-of-Speech tagging
func TestPOSTagging(t *testing.T) {
	client, err := sdk.NewClient()
	if err != nil {
		t.Fatalf("failed to initialize client: %v", err)
	}

	tests := []struct {
		name    string
		input   string
		wantErr bool
	}{
		{
			name:    "noun and verb",
			input:   "The cat runs",
			wantErr: false,
		},
		{
			name:    "adjective and noun",
			input:   "beautiful flower",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tags, err := client.POSTag(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("POSTag() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr && tags == nil {
				t.Error("expected non-nil tags")
			}
		})
	}
}

// TestNamedEntityRecognition tests NER functionality
func TestNamedEntityRecognition(t *testing.T) {
	client, err := sdk.NewClient()
	if err != nil {
		t.Fatalf("failed to initialize client: %v", err)
	}

	tests := []struct {
		name    string
		input   string
		wantErr bool
	}{
		{
			name:    "person and location",
			input:   "John Smith works in New York",
			wantErr: false,
		},
		{
			name:    "organization",
			input:   "Apple Inc. is headquartered in California",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			entities, err := client.ExtractEntities(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("ExtractEntities() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr && entities == nil {
				t.Error("expected non-nil entities")
			}
		})
	}
}

// TestSentimentAnalysis tests sentiment analysis
func TestSentimentAnalysis(t *testing.T) {
	client, err := sdk.NewClient()
	if err != nil {
		t.Fatalf("failed to initialize client: %v", err)
	}

	tests := []struct {
		name    string
		input   string
		wantErr bool
	}{
		{
			name:    "positive sentiment",
			input:   "I love this product!",
			wantErr: false,
		},
		{
			name:    "negative sentiment",
			input:   "This is terrible",
			wantErr: false,
		},
		{
			name:    "neutral sentiment",
			input:   "The weather is cloudy",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sentiment, err := client.AnalyzeSentiment(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("AnalyzeSentiment() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr && sentiment == nil {
				t.Error("expected non-nil sentiment")
			}
		})
	}
}

// TestLemmatization tests lemmatization functionality
func TestLemmatization(t *testing.T) {
	client, err := sdk.NewClient()
	if err != nil {
		t.Fatalf("failed to initialize client: %v", err)
	}

	tests := []struct {
		name    string
		input   string
		wantErr bool
	}{
		{
			name:    "plural to singular",
			input:   "running",
			wantErr: false,
		},
		{
			name:    "past tense",
			input:   "jumped",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lemmas, err := client.Lemmatize(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Lemmatize() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr && lemmas == nil {
				t.Error("expected non-nil lemmas")
			}
		})
	}
}

// TestDependencyParsing tests dependency parsing
func TestDependencyParsing(t *testing.T) {
	client, err := sdk.NewClient()
	if err != nil {
		t.Fatalf("failed to initialize client: %v", err)
	}

	input := "The quick brown fox jumps over the lazy dog"

	deps, err := client.ParseDependencies(input)
	if err != nil {
		t.Fatalf("ParseDependencies() error = %v", err)
	}

	if deps == nil {
		t.Error("expected non-nil dependencies")
	}
}

// TestSimilarity tests text similarity calculation
func TestSimilarity(t *testing.T) {
	client, err := sdk.NewClient()
	if err != nil {
		t.Fatalf("failed to initialize client: %v", err)
	}

	tests := []struct {
		name    string
		text1   string
		text2   string
		wantErr bool
	}{
		{
			name:    "identical texts",
			text1:   "hello world",
			text2:   "hello world",
			wantErr: false,
		},
		{
			name:    "similar texts",
			text1:   "the cat is here",
			text2:   "the dog is there",
			wantErr: false,
		},
		{
			name:    "different texts",
			text1:   "programming",
			text2:   "cooking",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			score, err := client.CalculateSimilarity(tt.text1, tt.text2)
			if (err != nil) != tt.wantErr {
				t.Errorf("CalculateSimilarity() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr && (score < 0 || score > 1) {
				t.Errorf("expected similarity score between 0 and 1, got %v", score)
			}
		})
	}
}

// TestBenchmarkTokenization benchmarks the tokenization performance
func BenchmarkTokenization(b *testing.B) {
	client, _ := sdk.NewClient()
	text := "The quick brown fox jumps over the lazy dog. This is a benchmark test."

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		client.Tokenize(text)
	}
}

// TestBenchmarkSentimentAnalysis benchmarks sentiment analysis performance
func BenchmarkSentimentAnalysis(b *testing.B) {
	client, _ := sdk.NewClient()
	text := "I absolutely love this product! It's amazing and works perfectly."

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		client.AnalyzeSentiment(text)
	}
}

// TestExample demonstrates basic SDK usage
func TestExample(t *testing.T) {
	// Initialize client
	client, err := sdk.NewClient()
	if err != nil {
		t.Fatalf("failed to initialize client: %v", err)
	}

	// Example text
	text := "ZenNLP is a powerful natural language processing library for Go."

	// Tokenize
	tokens, _ := client.Tokenize(text)
	fmt.Printf("Tokens: %v\n", tokens)

	// Analyze sentiment
	sentiment, _ := client.AnalyzeSentiment(text)
	fmt.Printf("Sentiment: %v\n", sentiment)

	// Extract entities
	entities, _ := client.ExtractEntities(text)
	fmt.Printf("Entities: %v\n", entities)

	// POS tagging
	tags, _ := client.POSTag(text)
	fmt.Printf("POS Tags: %v\n", tags)
}
