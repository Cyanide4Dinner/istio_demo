## Dev
- Build protobufs by `protoc --go_out=plugins=grpc:.`

## Building

- Run `eval $(minikube docker-env)` before for k8s to find image.
- Build by `docker build -t local/mult -f Dockerfile.mult .`

## Deploying
 
- Apply deployment `kubectl apply -f mult-deployment.yml`.
- Apply services `kubectl apply -f mult-deployment.yml`.

## Testing

- Get minikube URL for service by `minikube service --url grpc-istio-mult`.
- Trigger by `grpcurl -plaintext -d '{"a": 2, "b": 78}' 192.168.49.2:30228 pb.MultService/Compute`.

## Istio

- Apply AuthorizationPolicy using `kubectl apply -f allow_nothing.yml'.
- Remove by `kubectl delete authorizationpolicy.security.istio.io/allow-nothing`.

## AuthorizationPolicy
- `to` filter runs only while `when` is present even if it is empty for gRPC method filtering.

## BloomRPC
- Run BloomRPC with `./bloomrpc --disable-features=VizDisplayCompositor &'

## WireShark
- Go to 'any' interface.
- Add `grpc_person_search_protobuf_with_image.pcapng` with File > Open.
- For ease, filter the destination using `ip.dst == 192.168.49.2`.
- To view HTTP2/GRPC traffic, Analyze > Decode As, add TCP port as value of minikube port and enter 'HTTP2' in the Current column for the row.
- Show be able to view HTTP2 & GRPC traffic.
- 
