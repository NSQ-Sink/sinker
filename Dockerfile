# syntax=docker/dockerfile:1

FROM golang:1.20

WORKDIR /app

COPY . ./

EXPOSE 4150 4151 4160 4161 4170 4171

RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-nsq-sink

CMD ["/docker-nsq-sink", "--config-path", "example-config-podman.json"]