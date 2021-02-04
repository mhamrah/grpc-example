CONTAINER=gcr.io/grpc-demo-1

gen/todos.pb.go: todos.proto
.PHONY: gen
gen:
	docker run -v $(PWD):/workspace --rm grpckit/omniproto:1.35_0

.PHONY: protogen
protogen: gen/todos.pb.go

.PHONY: build
build: todos.proto
	docker build --target server -t ${CONTAINER}/todos-server .
	docker build --target client -t ${CONTAINER}/todos-client .

.PHONY: deploy-server
deploy-server: protogen
	docker push ${CONTAINER}/todos-server
	kubectl apply -f k8s/todos-server.yaml

.PHONY: deploy-client
deploy-client: protogen
	docker push ${CONTAINER}/todos-client
	kubectl apply -f k8s/todos-client.yaml

.PHONY: envoy
envoy:
	docker run -it --rm --name envoy -p 9901:9901 -p 9000:9000 -p 51051:51051 \
		-v `pwd`/gen/descriptors.pb:/data/descriptors.pb:ro \
		-v `pwd`/envoy-config.yml:/etc/envoy/envoy.yaml:ro \
		envoyproxy/envoy:v1.16-latest
