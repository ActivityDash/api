FROM golang:1.18.3-alpine

RUN addgroup -S docker && adduser -S -G docker docker \
    && mkdir -p /home/docker/Downloads \
    && chown -R docker:docker /home/docker

RUN mkdir -p /home/docker/app
WORKDIR /home/docker/app
COPY . /home/docker/app
RUN chown -R docker:docker /home/docker/app

RUN go mod download
RUN go build -o /api

# Run everything after as non-privileged user.
USER docker

EXPOSE 8080
CMD ["/api"]
