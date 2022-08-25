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

    docker container run -d ubuntu /bin/bash -c 'sleep 10'

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

Network
---
Exposing container ports on the host.

Set the working directory to /images/api-letters.  
Writing the Dockerfile.

    FROM golang:1.19.0
    COPY . /src
    WORKDIR /src
    RUN go build main.go
    EXPOSE 8080
    CMD ["./main"]

Running the container.

    docker image build -t api-letters .
    docker container run -p 5000:8080 api-letters

Connecting through the host.

    curl localhost:5000/letters

Composing containers
---
Managing multiple containers as a unit.

Set the working directory to /compose/multiple-containers.  
Writing the docker-compose.yml.

    version: "3.9"
    services:
      api-letters:
        build: ../../images/api-letters
        ports:
          - 5000:8080
      api-numbers:
        build: ../../images/api-numbers
        ports:
          - 5001:8081

Running the containers.

    docker-compose up
    docker-compose up -d

Inspecting the containers.

    docker-compose ps
    docker-compose top

Stopping the containers

    docker-compose stop

Restarting the containers.

    docker-compose restart

Deleting the containers.

    docker-compose stop
    docker-compose rm

    docker-compose down





Connecting containers
---
Set the working directory to /compose/network.  
Writing the docker-compose.yml.

    version: "3.9"
    services:
      api-sum:
        build: ../../images/api-sum
        environment:
          - API_NUMBERS_URL=http://api-numbers:8080
        ports:
          - 5001:8081
      api-numbers:
        build: ../../images/api-numbers
        expose:
          - 8080

Running the containers.

    docker-compose up
    curl localhost:5001/numbers
