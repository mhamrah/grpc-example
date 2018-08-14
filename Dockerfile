FROM golang:1.10-alpine3.8 as build

RUN apk add --update make git

WORKDIR /go/src/github.com/mhamrah/todos
COPY . .

RUN make deps

RUN go build -o bin/server server/cmd/main.go
RUN go build -o bin/client client/cmd/main.go

FROM alpine:3.8 as server

WORKDIR /app

COPY --from=build /go/src/github.com/mhamrah/todos/bin/server /app/server

EXPOSE 50051/tcp

ENTRYPOINT ["/app/server"]


FROM alpine:3.8 as client

WORKDIR /app

COPY --from=build /go/src/github.com/mhamrah/todos/bin/client /app/client

ENTRYPOINT ["/app/client"]
