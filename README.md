# The original code from my `gRPC All The Things` talk at Full Stack Fest 2018 is on the branch _original-fullstackfest-2018_. The master branch

tracks an upgraded project with the similar code.

[Slides](https://speakerdeck.com/mhamrah/grpc-all-the-things)

# TODOS

A grpc-based [todo backend](http://todobackend.com/) used to demonstrate [gRPC Gateway](https://github.com/grpc-ecosystem/grpc-gateway)
and running microservices on Kubernetes. This is demo code used for

## Running Locally

Generate the required protobuf files with `make protogen`. This will run
the `protoc` command in the `namely/protoc-all` docker container.
