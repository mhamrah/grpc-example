# TODOS

A grpc-based [todo backend](http://todobackend.com/) used to demonstrate [gRPC Gateway](https://github.com/grpc-ecosystem/grpc-gateway)
and running microservices on Kubernetes.

## Running Locally

Generate the required protobuf files with `make protogen`. This will run
the `protoc` command in the `namely/protoc-all` docker container.

You can use docker-compose to run the client, server and gateway container
locally. This dockr compose file also creates a default network of `grpc-example` you can later attach too for running `namely/grpc-cli` commands.

You can run with vanilla docker by running `make build` to build docker containers for the gateway,
client and server.

Alternatively run `make deps` to install govendor and sync package dependencies.
You can then start a server with `go run server/cmd/main.go`. You can
also run `go run client/cmd/main.go` to generate some synthetic load.

## Containers

This project has four docker-compose services that are built from
corresponding dockerfiles.

* `todos` is the backend service (go)
* `todos-gw` is the generated grpc-gateway from the todos.proto (go)
* `client` is a simple service that periodically calls the backend, to 
*   generate sythentic load. (go)
* `todos-graphq` is a rejoiner-based java app that wrapps the todos endpoint (java)

You can call these services from their name by connecting a docker container to the `grpc-example` 
network, like so:

```
docker run --network grpc-example -it --rm alpine /bin/sh
```

## Setting Up Kubernetes

This project has a terraform module to bring up a Kubernetes cluster. Once you've installed
Terraform and set up your GCP credentials using a service account, `cd` into the terraform
directory and run `terraform init`. You should create a `terraform.tfvars` file with your
project name, credential file location, username and password and run `terraform apply`. 
This brings up a 3-node Kubernetes cluster using f1-micro instance.

With your terraform cluster established configure kubectl:

`gcloud container clusters get-credentials grpc-demo-cluster --region us-east1`

Update the GCP project name in the Makefile, endpoints config, and k8s manifest (you can
find and replace `grpc-demo-1` with your project name). Run `make deploy-server`,
`make deploy-gateway`, and `make deploy-client` to build the appropriate containers and
deploy to Kubernetes. You can also run `make deploy-endpoints` to configure Cloud Endpoints.

