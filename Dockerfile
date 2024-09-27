FROM golang:1.23.1 AS builder
ARG CGO_ENABLED=0
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -ldflags="-w -s" -o=./bin/live-cursors.tmp ./cmd

FROM gruebel/upx:latest AS upx
ARG UPX_VERSION=4.2.4
WORKDIR /app
COPY --from=builder /app/bin .
COPY --from=builder /app/env.yml .
COPY --from=builder /app/banner.txt* .
RUN upx --lzma --best -o ./live-cursors ./live-cursors.tmp
RUN rm -rf ./live-cursors.tmp

FROM gcr.io/distroless/static
WORKDIR /app
COPY --from=upx /app .
EXPOSE 8080
ENTRYPOINT ["./live-cursors"]
