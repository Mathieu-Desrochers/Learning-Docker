Managing images
---
Searching for images.  
This is the official registry.

    https://hub.docker.com

Downloading images.

    docker pull redis
    docker pull redis:7.0
    docker pull redis:latest

Listing images.

    docker image ls

Inspecting images.

    docker image inspect redis

Deleting images.

    docker image rm redis

Managing containers
---
Starting a container with a specific command.  
The container stops when the command terminates.  
Pressing Ctrl-PQ detaches the terminal.

    docker container run -it redis /bin/bash

Starting a container as a deamon.  
With the image's default command or explicitly.

    docker container run -d redis
    docker container run -d redis /bin/sh -c redis-server

Listing containers.

    docker container ls
    docker container ls -a

Inspecting containers.

    docker container inspect bf59d85536e8

Sending commands to a running container.

    docker exec -it bf59d85536e8 /bin/bash

Stopping a container.  
The container is not deleted.

    docker container stop bf59d85536e8

Restarting a container.

    docker container start bf59d85536e8

Deleting a container.

    docker container stop bf59d85536e8
    docker container rm bf59d85536e8

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

Building images in multiple stages
---
Building an image stripped from the compiler and tools.  
Just the runtime requirements.

Set the working directory to /images/hello-world-stages.  
Writing the Dockerfile.

    FROM golang:1.19.0-bullseye AS compile
    COPY . /src
    WORKDIR /src
    RUN go build main.go

    FROM debian:bullseye-slim
    WORKDIR /usr/bin
    COPY --from=compile /src/main .
    CMD ["./main"]

The previous image was chunky at 994MB.  
We got it down to 82MB.

Composing applications
---
