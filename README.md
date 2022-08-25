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

    docker container run ubuntu
    docker container run ubuntu /bin/bash

Starting a container in the background.

    docker container run -d ubuntu /bin/sh -c 'sleep 10'

Listing containers.

    docker container ls
    docker container ls -a

Executing commands on a running container.

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

    docker image build -t image-hello-world:latest .

Running the image.

    docker container run image-hello-world:latest

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

    docker image build -t stages:latest .

The previous image was chunky at 994MB.  
We got it down to 82 MB.

Network
---
Exposing a container port on the host.

Set the working directory to /images/api-numbers.  
Writing the Dockerfile.

    FROM golang:1.19.0
    COPY . /src
    WORKDIR /src
    RUN go build main.go
    EXPOSE 8080
    CMD ["./main"]

Running the image.

    docker image build -t api-numbers:latest .
    docker container run -p 5000:8080 api-numbers:latest

Connecting through the host.

    curl localhost:5000/numbers

Compose
---
