FROM golang:1.20.3

RUN go version
ENV GOPATH=/

WORKDIR /usr/app

# install psql
RUN apt-get update && \
    apt-get -y install postgresql-client \
    htop

# install migrate
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.0/migrate.linux-amd64.tar.gz | tar xz -C /tmp/

RUN mv /tmp/migrate /usr/local/bin/migrate

RUN chmod +x /usr/local/bin/migrate

RUN rm -rf /tmp/migrate

COPY ./ ./

# build go app
RUN go mod download && \
    go build -o todo-go-app ./cmd/main.go

ENTRYPOINT ["./todo-go-app"]

CMD ["watch", "-n", "5", "ls"]