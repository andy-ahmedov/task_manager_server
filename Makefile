all: run


sup:
	sudo docker-compose up --build server

server_run: 
	go run server/cmd/main.go

client_run: 
ifeq ($(CMD),update)
	go run client/cmd/main.go $(CMD) $(ID) --name $(NAME) --description $(DESCRIPTION) --status $(STATUS)
else
	go run client/cmd/main.go $(CMD) $(filter-out $@,$(MAKECMDGOALS))
endif

# Exampe:
#	-$> make client_run create "Name" "Description" "Status"

# Example for using update command: 
#	-$> CMD=update ID=7 NAME="SomeName" DESCRIPTION='SomeDescriptio' STATUS='SomeStatus' make client_run

#	Important! When using the update command via make, 
#	only 1 word can be passed in parameters. If you want 
#	to use a command and pass several words to a specific 
#	field, then use without make

#	-$> go run client/cmd/main.go update 7 --status "Some Status"
%:
	@:

up:
	docker-compose up -d db

stop_and_delete_container:
	docker stop task_manager
	docker rm task_manager
	docker image rmi task_management_service-db:latest

create_table:
	docker exec -it task_manager psql -U postgres -d task_service -c "\i script.sql"

proto:
	protoc -I service_api/proto --go_out=service_api/api service_api/proto/task_manager.proto
	protoc -I service_api/proto --go-grpc_out=service_api/api service_api/proto/task_manager.proto
