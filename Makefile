all: run

run:
	sudo docker-compose up server
#	sudo docker-compose run server

build_up:
	sudo docker-compose up --build server

server_run: 
	go run cmd/main.go

up:
	docker-compose up -d db

stop_and_delete_container:
	docker stop new_task_manager
	docker rm new_task_manager
	docker image rmi task_manager_server-db:latest

create_table:
	docker exec -it new_task_manager psql -U postgres -d task_service -c "\i script.sql"