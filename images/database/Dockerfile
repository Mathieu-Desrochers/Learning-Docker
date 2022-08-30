FROM golang:1.19.0
COPY . /src
WORKDIR /src
RUN go build main.go
EXPOSE 8081
CMD ["./main"]
