# Build Container
FROM golang:1.17-alpine as builder

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN go build -o layered-architecture-demo ./src/cmd


# Runtime Container
FROM alpine:3.14

WORKDIR /

COPY --from=builder /app/layered-architecture-demo .

ENTRYPOINT ["/layered-architecture-demo"]