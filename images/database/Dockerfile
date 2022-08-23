FROM golang:1.19.0-bullseye AS compile
COPY . /src
WORKDIR /src
RUN go build main.go

FROM debian:bullseye-slim
WORKDIR /usr/local/bin
COPY --from=compile /src/main .
EXPOSE 8080
CMD ["./main"]
