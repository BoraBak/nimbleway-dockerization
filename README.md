# Web Server Container with Sequential Numbers and NGINX Load Balancing

This project sets up multiple web server containers using Docker Compose.  
Each web server displays a sequential number representing its startup order relative to the other containers.  
NGINX is used as a reverse proxy to load balance incoming requests among the web server containers.

## Implementation Details

### Docker Compose Configuration

- The `docker-compose.yaml` file defines three services: `nginx`, `web_sever`, and `sequence_generator`.
- NGINX service listens on port 80 and proxies requests to the `web_sever` service.
- The `web_sever` service runs the GOlang web server application, exposing port 8080.
- The `sequence_generator` service runs a GOlang application to generate sequential numbers, exposing port 9090.
- All services are connected to the same Docker network for internal communication.

### NGINX Configuration

- The NGINX configuration file (`nginx.conf`) sets up a reverse proxy to distribute incoming requests among the web server containers.
- NGINX listens on port 80 and forwards requests to the `web_sever` service.

### Web Server Application

- The web server application (`./web-server/main.go`) displays the server's name (set by an environment variable) along with a sequential number.
- The server name is provided by the `SERVER_NAME` environment variable.
- By default, if the `SERVER_NAME` variable is not provided, the system sets it to a default value of "web_server".
- The sequential number is retrieved from the `sequence_generator` service.

### Sequential Number Generator

- The sequential number generator (`./sequence-generator/main.go`) runs a simple HTTP server that increments a counter for each incoming request.
- The counter is reset when it reaches a specified sequence length.

### Dynamic Control Over Server Instances

- The number of web server instances is controlled dynamically using the `SEQUENCE_LENGTH` environment variable.
- This environment variable allows specifying the maximum number of server instances to be created.
- By default, if the `SEQUENCE_LENGTH` variable is not provided, the system sets it to a default value of 5.
- The `SEQUENCE_LENGTH` variable enables flexibility in scaling the number of server instances based on requirements without modifying the Docker Compose configuration.

## How to Run

1. Make sure you have Docker and Docker Compose installed on your machine.
2. Clone this repository to your local machine.
3. Navigate to the project directory.
4. Run `docker-compose up --build` to build and start the containers.
5. Access the web servers through the NGINX load balancer at `http://localhost`.

## Alternatives and Tradeoffs

- Instead of NGINX, other reverse proxies like HAProxy or Traefik could be used for load balancing.
- We could implement dynamic service discovery mechanisms like Consul or etcd for better scalability and fault tolerance.
- Using environment variables for server configuration may not be suitable for production-scale applications. Consider using configuration management tools or external configuration services.
- **Dynamic Scaling with Docker Compose:** Docker Compose allows us to scale services dynamically using the `docker-compose up --scale` command. We can specify the desired number of instances for a particular service using this command-line option. For example:
    ```bash
    docker-compose up --scale web_server=3
    ```
  This command will launch three instances of the `web_server` service.
- In a production environment, consider adding health checks and monitoring to ensure the availability and performance of the web servers and load balancer.

## Additional Features and Improvements

- Implementing SSL/TLS termination at the NGINX level for secure communication.
- Adding support for container orchestration platforms like Kubernetes for better scalability and management.
- Implementing automatic scaling based on traffic patterns using tools like Docker Swarm or Kubernetes.
- Integrating with centralized logging and monitoring solutions for better visibility into the system's performance and health.

## Rationale

- Docker Compose was chosen for container orchestration as it provides an easy way to define and manage multi-container applications.
- NGINX was chosen as the reverse proxy for its high performance, reliability, and extensive features for load balancing and proxying.
- The solution provides a simple yet effective way to set up multiple web server containers with sequential numbers and achieve load balancing for incoming requests.

