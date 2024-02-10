FROM golang:latest

# COPY server /github.com/andy-ahmedov/task_management_service/server/
# COPY service_api /github.com/andy-ahmedov/task_management_service/service_api/
# COPY go.mod /github.com/andy-ahmedov/task_management_service/
# COPY .env /github.com/andy-ahmedov/task_management_service/

COPY . /github.com/andy-ahmedov/task_manager_server/

WORKDIR /github.com/andy-ahmedov/task_manager_server/

RUN go mod download
RUN go build -o ./bin/server_init server/cmd/main.go

CMD ["./bin/server_init"]