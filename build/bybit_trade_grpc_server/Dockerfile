FROM golang:1.23.3-alpine AS builder

WORKDIR /build

COPY . .

ENV GO111MODULE=on

ENV GOCACHE=/root/.cache/go-build

RUN --mount=type=cache,target=/go/pkg/mod/ \
    --mount=type=bind,source=go.sum,target=go.sum \
    --mount=type=bind,source=go.mod,target=go.mod \
    go mod download -x

ENV GOCACHE=/root/.cache/go-build

RUN --mount=type=cache,target=/go/pkg/mod/ \
    --mount=type=cache,target="/root/.cache/go-build" \
    --mount=type=bind,target=. \
    go build -v -o  cmd/bybit/bybit_trade_grpc_server


FROM alpine:latest

WORKDIR /api/
ENV PATH=/api/bin/:$PATH

COPY --from=builder /build ./bin/bybit_trade_grpc_server
COPY --from=builder /build/env .


EXPOSE 9000

CMD ["/cmd/bybit/bybit_trade_grpc_server", "-env", "env"]
