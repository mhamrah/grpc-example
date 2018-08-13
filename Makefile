
gen/todo.pb.go: todo.proto
	docker run -v `pwd`:/defs namely/protoc-all:1.15 -f todo.proto -l gogo -o gen

.PHONY: compile
compile: gen/todo.pb.go

.PHONY: deps
deps:
	go get -u github.com/kardianos/govendor
	govendor sync
