CONTAINER=gcr.io/grpc-demo-1

gen/todos.pb.go: todos.proto
	docker run -v `pwd`:/defs namely/protoc-all:1.14_0 -f todos.proto -l go -o gen

.PHONY: protogen
protogen: gen/todos.pb.go

.PHONY: deps
deps:
	go get -u github.com/kardianos/govendor
	govendor sync

.PHONY: gateway
gateway: todos.proto
	docker run -v `pwd`/todos.proto:/defs/todos.proto -v `pwd`/gen:/defs/gen namely/gen-grpc-gateway -f todos.proto -s Todos
	docker build -t ${CONTAINER}/todos-gateway gen/grpc-gateway

.PHONY: deploy-gateway
deploy-gateway: gateway
	docker push ${CONTAINER}/todos-gateway
	kubectl apply -f k8s/todos-gateway.yaml

.PHONY: deploy-server
deploy-server: protogen
	docker build --target server -t ${CONTAINER}/todos-server .
	docker push ${CONTAINER}/todos-server
	kubectl apply -f k8s/todos-server.yaml

.PHONY: deploy-client
deploy-client: protogen
	docker build --target client -t ${CONTAINER}/todos-client .
	docker push ${CONTAINER}/todos-client
	kubectl apply -f k8s/todos-client.yaml

.PHONY: deploy-endpoints
deploy-endpoints:
	docker run -v `pwd`/todos.proto:/defs/todos.proto -v `pwd`/gen:/defs/gen namely/protoc:1.14_0 -I . --descriptor_set_out=gen/api_descriptor.pb --include_imports todos.proto
	gcloud endpoints services deploy gen/api_descriptor.pb gcp-endpoints/api_config.yaml
