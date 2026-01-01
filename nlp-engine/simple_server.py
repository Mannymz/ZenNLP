import grpc
import logging
from concurrent import futures
import time

import nlp_pb2
import nlp_pb2_grpc

class NLPManagerServicer(nlp_pb2_grpc.NLPManagerServicer):
    def AnalyzeSentiment(self, request, context):
        logging.info(f"Analyzing sentiment for text: '{request.text}' in language: '{request.lang}'")
        
        # Simple mock analysis for testing
        text_lower = request.text.lower()
        
        if any(word in text_lower for word in ['عالی', 'خوب', 'راضی', 'عالی است']):
            label = "positive"
            score = 0.85
        elif any(word in text_lower for word in ['بد', 'ناامید', 'توصیه نمی‌کنم']):
            label = "negative"
            score = 0.15
        else:
            label = "neutral"
            score = 0.50
        
        return nlp_pb2.SentimentResponse(label=label, score=score)

def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    nlp_pb2_grpc.add_NLPManagerServicer_to_server(NLPManagerServicer(), server)
    
    server.add_insecure_port('localhost:50051')
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
