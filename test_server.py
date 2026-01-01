import sys
import os
sys.path.insert(0, os.path.join(os.path.dirname(__file__), 'nlp-engine'))

import grpc
import nlp_pb2
import nlp_pb2_grpc

def test_sentiment():
    # Create a gRPC channel
    with grpc.insecure_channel('localhost:50051') as channel:
        stub = nlp_pb2_grpc.NLPManagerStub(channel)

        # Test texts
        texts = [
            "این محصول عالی است",
            "کیفیت بسیار بد بود",
            "قیمت مناسبی دارد"
        ]

        for text in texts:
            request = nlp_pb2.SentimentRequest(text=text, lang="fa")
            response = stub.AnalyzeSentiment(request)
            print(f"Text: {text}")
            print(f"Sentiment: {response.label}, Score: {response.score}")
            print("---")

if __name__ == '__main__':
    test_sentiment()