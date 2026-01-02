# ZenNLP Setup Instructions

## Current Status

✅ **Python Dependencies**: Installed successfully
- grpcio, grpcio-tools, transformers, torch, numpy

✅ **Python Protobuf Files**: Generated successfully
- `nlp_pb2.py` and `nlp_pb2_grpc.py` created in `nlp-engine/`

✅ **Go Modules**: Configured
- api/, go-sdk/, and examples/ modules set up

✅ **Core Implementation**: Complete
- Real ParsBERT model integration
- Idiomatic Go SDK with retries and timeouts
- Docker configuration

## Quick Start

### 1. Start the Python NLP Server
```bash
# From the project root
python nlp-engine/server.py
```

### 2. Run Go Client Example
Navigate to the `examples` directory and run `demo.go`:

```bash
cd examples
go run demo.go
```

## Docker Alternative (Recommended)

```bash
# Start the NLP engine
docker-compose up nlp-engine

# In another terminal, test the client
cd examples
go run demo.go
```

## Architecture

- **Python Engine**: ParsBERT model for Persian sentiment analysis (with `force_download=True` for robust model loading)
- **Go SDK**: High-performance client with automatic retries
- **gRPC**: Efficient communication protocol

## Features

- ✅ Real Persian sentiment analysis
- ✅ Automatic retry logic
- ✅ Timeout handling
- ✅ Type-safe protobuf communication
- ✅ Docker deployment ready
- ✅ Production-ready error handling

## Next Steps

The project is ready for use! The core functionality is implemented and working. You can:

1. Start the Python server
2. Use the Go SDK to analyze Persian text
3. Deploy with Docker for production use
