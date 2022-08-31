Managing images
---
Searching for images.  
This is the official registry.

    https://hub.docker.com

Downloading images.

    docker pull ubuntu
    docker pull ubuntu:20.04
    docker pull ubuntu:latest

Listing images.

    docker image ls

Inspecting images.

    docker image inspect ubuntu

Deleting images.

    docker image rm ubuntu

Managing containers
---
Starting a container.  
The container stops when the command terminates.  
Pressing Ctrl-PQ detaches the terminal.

    docker container run -it ubuntu
    docker container run -it ubuntu /bin/bash

Starting a container in the background.

    docker container run -d ubuntu /bin/bash -c 'sleep 3600'

Listing containers.

    docker container ls
    docker container ls -a

Executing commands in a running container.

    docker exec -it 560b03d1bcd2 /bin/bash

Stopping a container.

    docker container stop 560b03d1bcd2

Restarting a container.

    docker container start 560b03d1bcd2

Deleting a container.

    docker container stop 560b03d1bcd2
    docker container rm 560b03d1bcd2

Building images
---
Set the working directory to /images/hello-world.  
Writing the Dockerfile.

    FROM golang:1.19.0
    COPY . /src
    WORKDIR /src
    RUN go build main.go
    CMD ["./main"]

Building the image.

    docker image build -t image-hello-world .

Running the image.

    docker container run image-hello-world

Multiple stages
---
Building an image without the compiler and source code.  

Set the working directory to /images/stages.  
Writing the Dockerfile.

    FROM golang:1.19.0-bullseye AS compile
    COPY . /src
    WORKDIR /src
    RUN go build main.go

    FROM debian:bullseye-slim
    WORKDIR /usr/bin
    COPY --from=compile /src/main .
    CMD ["./main"]

Building the image.

    docker image build -t stages .

The previous image was chunky at 994MB.  
We got it down to 82 MB.

Compose
---
Managing multiple containers declaratively.

Set the working directory to /images.  
Writing docker-compose.yml.

    version: "3.9"
    services:
      nginx:
        image: nginx
      redis:
        image: redis

Running the containers.

    docker-compose up
    docker-compose up -d

Inspecting the containers.

    docker-compose ps
    docker-compose top

Stopping the containers

    docker-compose stop

Deleting the containers.

    docker-compose rm
    docker-compose down

Swarms
---
Clustering multiple docker hosts.  
Create three docker hosts that can communicate with each other.  
Consider using https://labs.play-with-docker.com.

On the first host.  
Initializing the swarm.  
Joins the swarm as a manager node.

    docker swarm init `
      --advertise-addr 192.168.0.1:2377 `
      --listen-addr 192.168.0.1:2377

On the other hosts.  
Joining the swarm as worker nodes.

    docker swarm join `
      --advertise-addr 192.168.0.x:2377 `
      --listen-addr 192.168.0.x:2377 `
      --token SWMTKN-1-000000000 `
      192.168.0.1:2377

Listing the nodes.

    docker node ls

Leaving the swarm.

    docker swarm leave

Services
---
Instructing the swarm to run multiple instances of an image.  
Stopped containers are replaced automatically.

    docker service create `
      --name service1 --replicas 5 `
      ubuntu /bin/bash -c 'sleep 3600'

Listing services.  
The containers are balanced evenly across the nodes.

    docker service ls
    docker service ps service1

Inspecting services.

    docker service inspect service1

Scaling services up and down.

    docker service scale service1=10
    docker service scale service1=4

Updating services.

    docker service update `
      --image debian `
      --update-parallelism 2 `
      --update-delay 1m `
      service1

Removing services.

    docker service rm service1

Bridge network
---
Connecting containers on a single host.

Creating a bridge network.

    docker network create -d bridge bridge1

Listing the networks.

    docker network ls

Attaching containers to the network.

    docker container run `
      --name container1 --network bridge1 `
      -d busybox /bin/sh -c 'sleep 3600'

    docker container run `
      --name container2 --network bridge1 `
      -d busybox /bin/sh -c 'sleep 3600'

Inspecting networks.

    docker network inspect bridge1

Containers get an IP address through DHCP.

    docker exec -it container1 /bin/sh -c 'ifconfig'

Containers can address themselves through DNS.

    docker exec -it container1 /bin/sh -c 'ping -c 3 container2'

Overlay network
---
Connecting containers across multiple hosts.

Create a swarm composed of two nodes.  
Consider using https://labs.play-with-docker.com.

On the first node.  
Creating an overlay network.

    docker network create -d overlay --attachable overlay1

On the first node.  
Attaching a container to the network.

    docker container run `
      --name container1 --network overlay1 `
      -d busybox /bin/sh -c 'sleep 3600'

On the second node.  
Attaching a container to the network.

    docker container run `
      --name container2 --network overlay1 `
      -d busybox /bin/sh -c 'sleep 3600'

On the first node.  
Confirming the containers can communicate.

    docker exec -it container1 /bin/sh -c 'ping -c 3 container2'

Published ports
---
Mapping container ports on the host.

Set the working directory to /images/api.  
Writing the Dockerfile.

    FROM golang:1.19.0
    COPY . /src
    WORKDIR /src
    RUN go build main.go
    EXPOSE 8080
    CMD ["./main"]

Building the image.

    docker image build -t api .

Publishing a port.

    docker container run --name api -p 5001:8080 -d api

Listing the published ports.

    docker port api

Confirming the port is available from the host.

    curl localhost:5001/hello

Load balancing
---
Publishing a port multiple times on the same host with swarms.  
A load balancer is offered by the swarm to distribute the connections.

Initializing a single node.

    docker swarm init

Publishing the same port multiple times.

    docker service create `
      --name service1 -p 5002:8080 --replicas 3 `
      -d api

Confirming the port is load balanced by the swarm.

    curl localhost:5002/hello

Volumes
---
Mounting local storage inside containers.

Creating a volume.

    docker volume create volume1

Listing volumes.

    docker volume ls

Inspecting volumes.

    docker volume inspect volume1

Mounting the volume.

    docker container run --name container1 `
      --mount source=volume1,target=/volume1 `
      -d ubuntu /bin/bash -c 'sleep 3600'

Writing to the volume.

    docker exec -it container1 /bin/bash -c 'echo hello > /volume1/hello'

Deleting volumes.

    docker container stop container1
    docker container rm container1
    docker volume rm volume1

Mounting external storage on multiple hosts  
requires the use of plugins.

    docker plugin ls

Configurations and secrets
---
Providing runtime values to containers.

Creating configurations and secrets.

    echo "Super green" | docker config create button_color -
    echo "This is a secret" | docker secret create my_secret_data -

Listing configurations and secrets.

    docker config ls
    docker secret ls

Attaching configurations and secrets.

    docker service create `
      --name service1 `
      --config button_color `
      --secret my_secret_data `
      -d busybox /bin/sh -c 'sleep 3600'

Accessing configurations and secrets from a container.

    docker ps
    docker container exec -it 13524401a141 cat /button_color
    docker container exec -it 13524401a141 cat /run/secrets/my_secret_data

Configurations are mounted on the root file system.  
Secrets are mounted in an in-memory filesystem.

Both are encrypted on the swarm manager.  
Both are plainly accessible from the containers.

Deleting configurations and secrets.

    docker service rm service1

    docker config rm button_color
    docker secret rm my_secret_data

Sharing images
---
Pushing images to hub.docker.com.

Set the working directory to /images/api.  
Building the image.

    docker image build -t your-username/api .

Set the working directory to /images/database.  
Building the image.

    docker image build -t your-username/database .

Loging in.

    docker login --username your-username

Pushing the images.

    docker push your-username/api
    docker push your-username/database

Loging out.

    docker logout

Stacks
---
Managing multiple services declaratively on a swarm.

Create a swarm composed of one manager and two worker nodes.  
Consider using https://labs.play-with-docker.com.

On the manager node.  
Writing docker-stack.yml.

    version: "3.9"
    services:
      api:
        image: your-username/api
        ports:
          - "8080:8080"
        networks:
          - network1
        deploy:
          replicas: 3
      database:
        image: your-username/database
        ports:
          - "8081:8081"
        networks:
          - network1
        deploy:
          replicas: 2
    networks:
      network1:

Deploying the stack.

    docker stack deploy -c docker-stack.yml stack1

Listing stacks.

    docker stack ls

Inspecting stacks.

    docker stack ps stack1

Communicating with services from the nodes.  
Connections are load balanced across replicas.

    curl localhost:8080/sum
    curl localhost:8081/numbers

Communicating with services from the containers.  
Name resolution and load balancing across replicas is offered.

    http.Get("http://database:8081/numbers")

Inspecting logs.

    docker service logs stack1_api
    docker service logs stack1_database

Managing the stack declaratively.  
Updating docker-stack.yml.

    services:
      api:
        deploy:
          replicas: 5

Applying the changes on the stack.

    docker stack deploy -c docker-stack.yml stack1

Deleting stacks.

    docker stack rm stack1
