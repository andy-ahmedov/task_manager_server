<h1 align="center">TASK MANAGER SERVER</h1>


<p align="center">
 <img alt="Golang" src="https://img.shields.io/badge/Golang-4285f4?style=for-the-badge&logo=go&labelColor=black">

 <img alt="PostgreSQL" src="https://img.shields.io/badge/PostgreSQL-%23212121?style=for-the-badge&logo=postgreSQL&labelColor=black">

 <img alt="RabbitMQ" src="https://img.shields.io/badge/RabbitMQ-%23ff6600?style=for-the-badge&logo=rabbitmq&labelColor=black">


  <img alt="MongoDB" src="https://img.shields.io/badge/MongoDB-%23011d2b?style=for-the-badge&logo=mongodb&labelColor=black">

 <img alt="Docker" src="https://img.shields.io/badge/Docker-%231d63ed?style=for-the-badge&logo=docker&labelColor=black">
</p>
<hr>

![server](https://i.imgur.com/vtXhJBU.png)


## ABOUT ##

This is a service written in Go that provides a user-friendly interface for managing users and books via a REST API. The service supports user registration and authentication using JWT tokens, ensuring security and efficiency. The service provides the ability to create, update, delete and obtain information about books, access to which is limited to authenticated users.

Swagger documentation is provided for the convenience of users and developers. Interaction with the audit-log-server is carried out via gRPC to write activity logs to the MongoDB database.

The service uses PostgreSQL to store information about users, books, and refresh tokens, and MongoDB to store logs. Both databases run in Docker containers for ease of deployment and scalability.

Final technology stack: Go, Gin, REST API, gRPC, PostgreSQL, MongoDB, Docker, JWT. This service is an excellent example of modern, secure and scalable web development.

<hr>

## ⚡ STARTING ##

```bash
# Clone this project
$ git clone https://github.com/andy-ahmedov/crud_service

# Create in the root directory of the project and fill in the .env file

# Up MongoDB
$ make mongodb

# Up PostgreSQL
$ make up

# Create Tables
$ make create_table

# Run the project
$ make run

```

## 🔨 STACK ##
  <a href="https://go.dev/" target="_blank"> <img src="https://github.com/andy-ahmedov/andy-ahmedov/blob/main/files/Go-Logo_Yellow.png?raw=true" alt="golang" align="left" height='50px'/> </a>
  <a href="https://www.gnu.org/software/bash/" target="_blank"> <img src="https://raw.githubusercontent.com/andy-ahmedov/README_icons/5a0bd0723991e5d95e0eb90ce4e544345b69e05b/language_and_tools/square/bash/bash.svg" alt="bash" align="left" height='50px'/> </a>
  <a href="https://www.docker.com/" target="_blank"> <img src="https://raw.githubusercontent.com/andy-ahmedov/README_icons/5a0bd0723991e5d95e0eb90ce4e544345b69e05b/language_and_tools/square/docker/docker.svg" alt="docker" align="left" height='50px'/> </a>
  <a href="https://www.rabbitmq.com/" target="_blank"> <img src="https://raw.githubusercontent.com/andy-ahmedov/README_icons/5a0bd0723991e5d95e0eb90ce4e544345b69e05b/language_and_tools/square/rabbitmq/rabbitmq.svg" align="left" alt="rabbitmq" height='50px'/> </a>
  <a href="https://www.mongodb.com/" target="_blank"> <img src="https://raw.githubusercontent.com/devicons/devicon/6910f0503efdd315c8f9b858234310c06e04d9c0/icons/mongodb/mongodb-plain-wordmark.svg" align="left" alt="mongodb" height='50px'/> </a>
  <a href="https://www.postgresql.org/" target="_blank"> <img src="https://raw.githubusercontent.com/devicons/devicon/6910f0503efdd315c8f9b858234310c06e04d9c0/icons/postgresql/postgresql-plain-wordmark.svg" align="left" alt="postgres" height='50px'/> </a>
  <a href="https://grpc.io/" target="_blank"> <img src="https://raw.githubusercontent.com/devicons/devicon/6910f0503efdd315c8f9b858234310c06e04d9c0/icons/grpc/grpc-original.svg" align="left" alt="grpc" height='50px'/> </a>
  <a href="https://www.json.org/json-en.html" target="_blank"> <img src="https://raw.githubusercontent.com/devicons/devicon/6910f0503efdd315c8f9b858234310c06e04d9c0/icons/json/json-plain.svg" align="left" alt="json" height='50px'/> </a>
  <a href="https://code.visualstudio.com/" target="_blank"> <img src="https://raw.githubusercontent.com/devicons/devicon/6910f0503efdd315c8f9b858234310c06e04d9c0/icons/visualstudio/visualstudio-plain-wordmark.svg" align="left" alt="vscode" height='50px'/> </a>
