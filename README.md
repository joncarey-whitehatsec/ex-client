# Example Client

Example gRPC client for data-service-discovered-vuln. When the data service is running in okteto, this will establish a connection, send a couple of requests, and log the result to stderr.

```sh
$ go run cmd/ex-client/main.go
6:01PM INF unary call finished error="rpc error: code = InvalidArgument desc = required field attack_vector missing" grpc.code=InvalidArgument grpc.method=AddAttackVector grpc.service=discoveredvulnservice.DiscoveredVulnService grpc.start_time=2020-09-17T18:01:51-05:00 grpc.target=localhost:12001 grpc.time_ms=2.217816 span.kind=client
6:01PM ERR add av response error="rpc error: code = InvalidArgument desc = required field attack_vector missing" res=null
6:01PM INF send message finished grpc.code=OK grpc.method=ListAttackVectors grpc.service=discoveredvulnservice.DiscoveredVulnService grpc.start_time=2020-09-17T18:01:51-05:00 grpc.target=localhost:12001 grpc.time_ms=0.009762 span.kind=client
6:01PM ERR list av response stream={"ClientStream":{"ClientStream":{}}}
6:01PM ERR receive message finished error="rpc error: code = Unimplemented desc = method ListAttackVectors not implemented" grpc.code=Unimplemented grpc.method=ListAttackVectors grpc.service=discoveredvulnservice.DiscoveredVulnService grpc.start_time=2020-09-17T18:01:51-05:00 grpc.target=localhost:12001 grpc.time_ms=3.47955 span.kind=client
6:01PM ERR list av message error="rpc error: code = Unimplemented desc = method ListAttackVectors not implemented" msg=null
```
