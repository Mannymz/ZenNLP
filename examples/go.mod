module github.com/Mannymz/ZenNLP/examples

go 1.21

require (
	github.com/zen-nlp/go-sdk v0.0.0
	google.golang.org/grpc v1.60.0
)

replace github.com/zen-nlp/go-sdk => ../go-sdk

replace github.com/zen-nlp/api => ../api

require (
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/zen-nlp/api v0.0.0 // indirect
	golang.org/x/net v0.17.0 // indirect
	golang.org/x/sys v0.13.0 // indirect
	golang.org/x/text v0.13.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20231016165738-49dd2c1f3d0b // indirect
	google.golang.org/protobuf v1.31.0 // indirect
)