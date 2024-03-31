# syntax=docker/dockerfile:1
# how to run dockerfile, 
# example:
# 1. docker build .
# 2. docker build --build-arg WASHTUB="127.0.0.1:9000" --build-arg CONFIG_PATH="./example-config.json" --no-cache .

FROM golang:1.19

# get args
ARG WASHTUB="127.0.0.1:9000"
ARG CONFIG_PATH="./example-config.json"

# set workdir
WORKDIR /app

# copy all files
COPY . /

# get dependencies
RUN go mod download

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /sink

# Run
CMD ["/sink -washtub=$WASHTUB -config-path=$CONFIG_PATH"]


