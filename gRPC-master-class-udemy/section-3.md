## Types

1. Unary
2. server Streaming
3. client streaming
4. bidirectional streaming

- strict schema
- bidirectional


protoc -I hello/proto --go_out=.  --go_opt=module=github.com/oshanz/grpc-hello --go-grpc_out=.  --go-grpc_opt=module=github.com/oshanz/grpc-hello hello/proto/dummy.proto
