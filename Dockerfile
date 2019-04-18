FROM golang:latest AS builder
workdir /src
ADD ./app /src
RUN go get github.com/eclipse/paho.mqtt.golang && \
    go get github.com/prometheus/client_golang/prometheus && \
    go get github.com/prometheus/client_golang/prometheus/promhttp && \
    go get github.com/prometheus/common/log && \
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o /main .

# final stage
FROM alpine:latest
WORKDIR /app
COPY --from=builder /main /app/app
ENTRYPOINT ./app
