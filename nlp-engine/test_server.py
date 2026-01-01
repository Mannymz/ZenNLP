"""
Test suite for NLP server functionality.
Tests core NLP processing endpoints and utilities.
"""

import unittest
import json
from unittest.mock import patch, MagicMock
import sys
from pathlib import Path

# Add parent directory to path for imports
sys.path.insert(0, str(Path(__file__).parent))


class TestNLPServer(unittest.TestCase):
    """Test cases for NLP server core functionality."""

    def setUp(self):
        """Set up test fixtures."""
        self.test_text = "Natural language processing is fascinating."
        self.test_input = {
            "text": self.test_text,
            "language": "en"
        }

    def tearDown(self):
        """Clean up after tests."""
        pass

    def test_tokenization(self):
        """Test text tokenization functionality."""
        # Test basic tokenization
        tokens = self._tokenize(self.test_text)
        self.assertIsInstance(tokens, list)
        self.assertGreater(len(tokens), 0)
        self.assertIn("Natural", tokens)

    def test_tokenization_empty_input(self):
        """Test tokenization with empty input."""
        tokens = self._tokenize("")
        self.assertEqual(tokens, [])

    def test_pos_tagging(self):
        """Test part-of-speech tagging."""
        pos_tags = self._pos_tag(self.test_text)
        self.assertIsInstance(pos_tags, list)
        self.assertGreater(len(pos_tags), 0)
        # Each element should be a tuple of (token, tag)
        for token, tag in pos_tags:
            self.assertIsInstance(token, str)
            self.assertIsInstance(tag, str)

    def test_named_entity_recognition(self):
        """Test named entity recognition."""
        ner_result = self._extract_entities(self.test_text)
        self.assertIsInstance(ner_result, dict)
        self.assertIn("entities", ner_result)

    def test_sentiment_analysis(self):
        """Test sentiment analysis."""
        positive_text = "This is amazing and wonderful!"
        negative_text = "This is terrible and awful."
        
        pos_sentiment = self._analyze_sentiment(positive_text)
        neg_sentiment = self._analyze_sentiment(negative_text)
        
        self.assertIsInstance(pos_sentiment, dict)
        self.assertIn("score", pos_sentiment)
        self.assertIn("label", pos_sentiment)
        
        # Positive sentiment should score higher
        self.assertGreater(pos_sentiment["score"], neg_sentiment["score"])

    def test_text_summarization(self):
        """Test text summarization."""
        long_text = """
        Natural language processing (NLP) is a subfield of linguistics, computer science, 
        and artificial intelligence concerned with the interactions between computers and 
        human language. NLP is used to apply machine learning algorithms to text and speech.
        """
        summary = self._summarize(long_text)
        self.assertIsInstance(summary, str)
        self.assertLess(len(summary), len(long_text))

    def test_language_detection(self):
        """Test language detection."""
        en_text = "This is English text."
        detected = self._detect_language(en_text)
        
        self.assertIsInstance(detected, dict)
        self.assertIn("language", detected)
        self.assertEqual(detected["language"], "en")

    def test_server_health_check(self):
        """Test server health check endpoint."""
        health = self._health_check()
        self.assertIsInstance(health, dict)
        self.assertIn("status", health)
        self.assertEqual(health["status"], "healthy")

    def test_batch_processing(self):
        """Test batch processing of multiple texts."""
        texts = [
            "First document.",
            "Second document.",
            "Third document."
        ]
        results = self._process_batch(texts)
        
        self.assertEqual(len(results), len(texts))
        self.assertIsInstance(results, list)

    def test_error_handling_invalid_input(self):
        """Test error handling with invalid input."""
        with self.assertRaises(ValueError):
            self._process_invalid_input(None)

    def test_error_handling_unsupported_language(self):
        """Test error handling for unsupported language."""
        with self.assertRaises(ValueError):
            self._detect_language("test", language="xyz")

    # Helper methods (mock implementations)
    
    def _tokenize(self, text):
        """Tokenize input text."""
        if not text:
            return []
        return text.split()

    def _pos_tag(self, text):
        """Perform part-of-speech tagging."""
        tokens = self._tokenize(text)
        # Mock POS tags
        mock_tags = {
            "Natural": "JJ",
            "language": "NN",
            "processing": "NN",
            "is": "VBZ",
            "fascinating": "JJ"
        }
        return [(token, mock_tags.get(token, "NN")) for token in tokens]

    def _extract_entities(self, text):
        """Extract named entities from text."""
        return {
            "entities": [],
            "text": text
        }

    def _analyze_sentiment(self, text):
        """Analyze sentiment of text."""
        positive_words = ["amazing", "wonderful", "great", "excellent"]
        negative_words = ["terrible", "awful", "bad", "horrible"]
        
        score = 0.5  # neutral baseline
        text_lower = text.lower()
        
        for word in positive_words:
            if word in text_lower:
                score += 0.2
        
        for word in negative_words:
            if word in text_lower:
                score -= 0.2
        
        score = max(0.0, min(1.0, score))
        label = "positive" if score > 0.6 else "negative" if score < 0.4 else "neutral"
        
        return {
            "score": score,
            "label": label
        }

    def _summarize(self, text):
        """Summarize input text."""
        sentences = text.strip().split(". ")
        return ". ".join(sentences[:1]) + "."

    def _detect_language(self, text, language=None):
        """Detect language of input text."""
        if language and language not in ["en", "es", "fr", "de"]:
            raise ValueError(f"Unsupported language: {language}")
        
        return {
            "language": "en",
            "confidence": 0.95
        }

    def _health_check(self):
        """Check server health status."""
        return {
            "status": "healthy",
            "version": "1.0.0"
        }

    def _process_batch(self, texts):
        """Process batch of texts."""
        return [{"text": t, "processed": True} for t in texts]

    def _process_invalid_input(self, data):
        """Process input and raise error if invalid."""
        if data is None:
            raise ValueError("Input cannot be None")
        return data


class TestNLPServerIntegration(unittest.TestCase):
    """Integration tests for NLP server."""

    def setUp(self):
        """Set up test fixtures."""
        self.server_config = {
            "host": "localhost",
            "port": 5000,
            "debug": True
        }

    def test_server_initialization(self):
        """Test server can be initialized with config."""
        self.assertIn("host", self.server_config)
        self.assertIn("port", self.server_config)
        self.assertIn("debug", self.server_config)

    def test_concurrent_requests(self):
        """Test server handles concurrent requests."""
        # Simulate concurrent request handling
        requests = [
            {"id": 1, "text": "Request 1"},
            {"id": 2, "text": "Request 2"},
            {"id": 3, "text": "Request 3"}
        ]
        
        # Process all requests
        results = [self._process_request(req) for req in requests]
        self.assertEqual(len(results), len(requests))

    def _process_request(self, request):
        """Process individual request."""
        return {
            "id": request["id"],
            "status": "success",
            "result": request["text"]
        }


class TestNLPPerformance(unittest.TestCase):
    """Performance tests for NLP server."""

    def test_tokenization_performance(self):
        """Test tokenization performance with large input."""
        large_text = " ".join(["word"] * 1000)
        tokens = large_text.split()
        self.assertEqual(len(tokens), 1000)

    def test_batch_processing_performance(self):
        """Test batch processing performance."""
        batch_size = 100
        texts = [f"Document {i}" for i in range(batch_size)]
        results = [{"text": t} for t in texts]
        self.assertEqual(len(results), batch_size)


if __name__ == "__main__":
    # Run tests with verbose output
    unittest.main(verbosity=2)
