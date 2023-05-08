## Go Microservice

This is a Golang microservices project built with Kubernetes, Helm, Docker, and Docker Compose. The project consists of three endpoints - /health, /, and /details. The /health endpoint returns a JSON response with the application's status and timestamp, the / endpoint serves the homepage with a simple message, and the /details endpoint fetches and returns the hostname and IP address of the server.

### Technologies Used
The project is built with the following technologies:

- **Golang**: Golang is a popular programming language used to build fast, scalable, and efficient applications. It is used as the primary language to develop this microservices project.
- **Kubernetes**: Kubernetes is an open-source container orchestration platform used to manage and deploy containerized applications. It is used to deploy and manage the microservices in this project.
- **Helm**: Helm is a package manager for Kubernetes used to manage and deploy applications. It is used to package, install, and manage the microservices in this project.
- **Docker**: Docker is a containerization platform used to package and distribute applications. It is used to containerize the microservices in this project.
- **Docker Compose**: Docker Compose is a tool used to define and run multi-container Docker applications. It is used to define and run the containers in this project.

### Features
The project contains the following features:

- A Golang microservices application with three endpoints - `/health`, `/`, and `/details`.
- Kubernetes manifests and Helm charts for deployment and management of the microservices.
- Dockerfiles to containerize the microservices.
- A Docker Compose file to define and run the containers.

### Running Commands
To run the project, follow these steps:

1. Clone the repository to your local machine.
2. Navigate to the root directory of the project.
3. Run docker-compose up to start the containers.
4. Access the microservices at localhost:8080.