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
The container stops when the commands terminates.  
Pressing Ctrl-PQ detaches the terminal.

    docker container run -it redis /bin/bash

Starting a container as a deamon.  
With the image's default command or explicitly.

    docker container run -d redis
    docker container run -d redis /bin/sh -c redis-server

Listing containers.

    docker container ls

Stopping a container.

    docker container stop 7919c1dda089
