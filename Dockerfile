FROM golang:latest


WORKDIR /build
COPY ./ ./
RUN go build -o main cmd/main.go
CMD ["./main"]
