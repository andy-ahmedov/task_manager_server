FROM golang:latest AS builder
COPY . /github.com/andy-ahmedov/task_manager_server/
WORKDIR /github.com/andy-ahmedov/task_manager_server/
RUN go mod download

# Instructions for running without further assembly:
# RUN go build -o ./bin/server_init cmd/main.go
# CMD ["./bin/server_init"]

RUN GOOS=linux go build -o ./.bin/server_init cmd/main.go

FROM alpine:3.19
WORKDIR /root/
COPY --from=builder /github.com/andy-ahmedov/task_manager_server/.bin/server_init .
RUN apk add libc6-compat
# Needed to support running the executable file in the alpine environment.
CMD ["./server_init"]
