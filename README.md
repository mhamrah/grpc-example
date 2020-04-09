# The original code from my `gRPC All The Things` talk at Full Stack Fest 2018 is on the branch _original-fullstackfest-2018_. The master branch

tracks an upgraded project with the similar code.

[Slides](https://speakerdeck.com/mhamrah/grpc-all-the-things)

# TODOS

A grpc-based [todo backend](http://todobackend.com/) used to demonstrate [gRPC Gateway](https://github.com/grpc-ecosystem/grpc-gateway)
and running microservices on Kubernetes. This is demo code used for

## Install

```
brew cask install bloomrpc
brew install grpcurl
brew install minikube
brew install istioctl
```

## Running Locally

Generate the required protobuf files with `make protogen`. This will run
the `protoc` command in the `namely/protoc-all` docker container.

## Kubernetes

This project has a skaffold.yaml file and can be run on a minikube-based environment. Install the istioctl and minikube dependencies to experiment then do the following:

1. Start minikube (check for cpu/mem requirements): `minikube start`
1. Install the demo profile of istio, with Grafana enabled: `istioctl manifest apply --set profile=demo --set addonComponents.grafana.enabled=true`
1. In a separate terminal window, use Minikube tunnel to easily access services: `minikube tunnel`
1. Apply the `todos` namespace definition: `k apply -f k8s/setup/namespace.yaml`
1. Set istio to auto-inject on `todos`: `kubectl label namespace todos istio-injection=enabled`
1. Run skaffold in dev mode: `skaffold dev`

## Commands

```
grpcurl -plaintext localhost:50052 describe
grpcurl -plaintext -d '{ "id": "01E4Q00M7YPD06TX3YW8DMFF8B" }' localhost:50052 todos.Todos/GetTodo
```

```
curl  -X POST http://localhost:51051/todos -H 'Content-Type: application/json' -d '{ "id":  "01E4Q00M7YPD06TX3YW8DMFF8B" }'
```
