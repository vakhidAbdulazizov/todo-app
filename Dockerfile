FROM golang:1.20.3

RUN apt-get update && apt-get -y install postgresql-client htop \
    && curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.0/migrate.linux-amd64.tar.gz | tar xz -C /usr/local/bin \
    && chmod +x /usr/local/bin/migrate \
    && rm -rf /var/lib/apt/lists/*

ENV GOPATH=/

WORKDIR /usr/app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o todo-go-app ./cmd/main.go

ENTRYPOINT ["./todo-go-app"]
CMD ["watch", "-n", "5", "ls"]
