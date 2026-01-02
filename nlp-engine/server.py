import sys
import os

sys.path.append(os.path.dirname(__file__))

import grpc
import logging
from concurrent import futures
import time
import torch
from transformers import AutoTokenizer, AutoModelForSequenceClassification
import numpy as np

import nlp_pb2
import nlp_pb2_grpc

class NLPManagerServicer(nlp_pb2_grpc.NLPManagerServicer):
    def __init__(self):
        logging.info("Loading ParsBERT model...")
        self.model_name = "HooshvareLab/bert-fa-base-uncased-sentiment-snappfood"
        self.tokenizer = AutoTokenizer.from_pretrained(self.model_name, force_download=True)
        self.model = AutoModelForSequenceClassification.from_pretrained(self.model_name, force_download=True)
        self.labels = ["negative", "positive"]
        logging.info("Model loaded successfully")
    
    def AnalyzeSentiment(self, request, context):
        logging.info(f"Analyzing sentiment for text: '{request.text}' in language: '{request.lang}'")
        
        try:
            # Tokenize input text
            inputs = self.tokenizer(
                request.text, 
                return_tensors="pt", 
                truncation=True, 
                padding=True, 
                max_length=512
            )
            
            # Get model predictions
            with torch.no_grad():
                outputs = self.model(**inputs)
                logits = outputs.logits
                probabilities = torch.softmax(logits, dim=-1)
                predicted_class = torch.argmax(probabilities, dim=-1).item()
                confidence = probabilities[0][predicted_class].item()
            
            # Map to label and score
            label = self.labels[predicted_class]
            score = float(confidence)
            
            logging.info(f"Predicted sentiment: {label} with confidence: {score:.4f}")
            
            return nlp_pb2.SentimentResponse(label=label, score=score)
            
        except Exception as e:
            logging.error(f"Error during sentiment analysis: {str(e)}")
            context.set_code(grpc.StatusCode.INTERNAL)
            context.set_details(f"Sentiment analysis failed: {str(e)}")
            return nlp_pb2.SentimentResponse(label="error", score=0.0)

def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    nlp_pb2_grpc.add_NLPManagerServicer_to_server(NLPManagerServicer(), server)
    
    server.add_insecure_port('[::]:50051')
    logging.info("Starting gRPC server on port 50051")
    
    server.start()
    try:
        while True:
            time.sleep(86400)  # One day
    except KeyboardInterrupt:
        server.stop(0)

if __name__ == '__main__':
    logging.basicConfig(level=logging.INFO)
    serve()
