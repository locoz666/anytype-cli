# syntax=docker/dockerfile:1.6
FROM --platform=$TARGETPLATFORM golang:1.24-bookworm AS builder

ARG TARGETOS
ARG TARGETARCH

WORKDIR /src

RUN apt-get update \
    && apt-get install -y --no-install-recommends \
        ca-certificates \
        curl \
        gcc \
        g++ \
        git \
        make \
    && rm -rf /var/lib/apt/lists/*

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN make build

FROM debian:bookworm-slim

RUN apt-get update \
    && apt-get install -y --no-install-recommends \
        ca-certificates \
        libgcc-s1 \
        libstdc++6 \
    && rm -rf /var/lib/apt/lists/*

RUN useradd -r -u 10001 -g 0 -s /usr/sbin/nologin anytype \
    && mkdir -p /data \
    && chown -R 10001:0 /data

ENV HOME=/data
ENV DATA_PATH=/data

WORKDIR /data

COPY --from=builder /src/dist/anytype /usr/local/bin/anytype

EXPOSE 31012

USER 10001

ENTRYPOINT ["anytype"]
CMD ["serve", "--listen-address", "0.0.0.0:31012"]
