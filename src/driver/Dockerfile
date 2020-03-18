FROM golang:1.14
WORKDIR /go/src/
COPY . .
RUN GOOS=linux go build driver.go
EXPOSE 8081
ENTRYPOINT ["./driver"]