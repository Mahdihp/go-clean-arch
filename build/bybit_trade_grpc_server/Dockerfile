FROM golang:1.23.0-alpine AS builder

# Add Maintainer Info
LABEL maintainer="Mahdihp <mahdihp.devsc@gmail.com>"

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy everything from the current directory to the Working Directory inside the container
COPY . .
RUN ls
RUN pwd
# Build the Go app
ENV CGO_ENABLED=0
ENV GOOS=linux


RUN ls ./cmd/bybit/bybit_trade_grpc_server
ENV GOCACHE=/root/.cache/go-build
#RUN --mount=type=cache,target="/root/.cache/go-build" go build -o main .
RUN go build -o main  ./cmd/bybit/bybit_trade_grpc_server

RUN ls
RUN pwd

######## Start a new stage from scratch #######
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/
RUN ls
RUN pwd
# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/main .
COPY --from=builder /app/.env .

EXPOSE 9000

# Command to run the executable
CMD ["./main"]