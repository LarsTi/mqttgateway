FROM golang:latest AS builder
workdir /src
RUN go get github.com/eclipse/paho.mqtt.golang && \
    go get github.com/prometheus/client_golang/prometheus && \
    go get github.com/prometheus/client_golang/prometheus/promhttp && \
    go get github.com/prometheus/common/log
Add ./app /src
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o /main .

# final stage
FROM alpine
WORKDIR /app
COPY --from=builder /main /app/app
ENTRYPOINT ./app
