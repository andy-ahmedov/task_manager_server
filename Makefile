all: run

run:
	sudo docker-compose up server
#	sudo docker-compose run server

build_up:
	sudo docker-compose up --build server

server_run: 
	go run cmd/main.go

up:
	docker-compose up -d postgresdb

stop_and_delete_container:
	docker stop new_task_manager
	docker rm new_task_manager
	docker image rmi task_manager_server-db:latest

mongodb:
	docker run --rm -d --name mongodb -e MONGO_INITDB_ROOT_USERNAME=admin -e MONGO_INITDB_ROOT_PASSWORD=admin -e MONGO_INITDB_DATABASE=log_item_data -p 27017:27017 mongo:latest --auth

container_rune:
	docker run -it task_manager_debug:latest sh

create_table:
	docker exec -it new_task_manager psql -U postgres -d task_service -c "\i script.sql"