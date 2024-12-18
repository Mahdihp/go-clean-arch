FROM golang:1.23.2 AS builder

WORKDIR /bybit_trde_grpc_server_app

COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /go-clean-arch

##
## Run the tests in the container
##

FROM build-stage AS bybit_trde_grpc_server_app-stage
RUN go test -v ./...

##
## Deploy the application binary into a lean image
##

FROM gcr.io/distroless/base-debian11 AS build-release-stage

WORKDIR /

COPY --from=build-stage /go-clean-arch /go-clean-arch

EXPOSE 9000

USER nonroot:nonroot

ENTRYPOINT ["/go-clean-archg"]