# Multi-stage build for ZenNLP
FROM python:3.9-slim as nlp-engine

WORKDIR /app

# Install system dependencies
RUN apt-get update && apt-get install -y \
    gcc \
    g++ \
    && rm -rf /var/lib/apt/lists/*

# Copy Python requirements and install dependencies
COPY nlp-engine/requirements.txt .
RUN pip install --no-cache-dir -r requirements.txt

# Copy protobuf files and generate Python bindings
COPY api/nlp.proto .
RUN python -m grpc_tools.protoc \
    --proto_path=. \
    --python_out=. \
    --grpc_python_out=. \
    nlp.proto

# Copy Python server code
COPY nlp-engine/server.py .

# Expose port
EXPOSE 50051

# Run the server
CMD ["python", "server.py"]
