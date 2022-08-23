FROM golang:1.19.0-bullseye AS compile
COPY . /src
WORKDIR /src
RUN go build main.go

FROM debian:bullseye-slim
WORKDIR /usr/bin
COPY --from=compile /src/main .
CMD ["./main"]
