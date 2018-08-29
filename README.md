# TODOS

A grpc-based [todo backend](http://todobackend.com/) used to demonstrate [gRPC Gateway](https://github.com/grpc-ecosystem/grpc-gateway)
and running microservices on Kubernetes.

## Setting Up

This project has a terraform module to bring up a Kubernetes cluster. Once you've installed
Terraform and set up your GCP credentials, `cd` into the terraform
directory and run `terraform apply`. This brings up a 3-node Kubernetes cluster using f1-micro
preemtible instances. 

With your terraform cluster established configure kubectl:

`gcloud container clusters get-credentials grpc-demo-cluster --region us-east1`

## Running Locally

Generate the required protobuf files with `make protogen`. This will run
the `protoc` command in the `namely/protoc-all` docker container.


Run `make deps` to install govendor and sync package dependencies.
You can then start a server with `go run server/cmd/main.go`. You can
also run `go run client/cmd/main.go` to generate some synthetic load.


