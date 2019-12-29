# Multistage image keeps image size very small 
FROM golang:1.12 AS builder
WORKDIR /app
COPY go.mod go.mod
COPY go.sum go.sum
RUN go mod download
COPY . /app
RUN go build -o server

FROM gcr.io/distroless/base-debian10
WORKDIR /app
COPY --from=builder /app/server /
CMD ["./server"]