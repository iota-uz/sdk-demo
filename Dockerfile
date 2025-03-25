FROM iotauz/sdk:base-0.0.1

WORKDIR /build

# Install PostgreSQL client
RUN apt-get update && apt-get install -y postgresql-client && apt-get clean

COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN make css

RUN go build -o run_server cmd/server/main.go && go build -o seed_db cmd/seed/main.go
RUN go build -o migrate cmd/migrate/main.go
RUN /build/migrate collect
RUN chmod +x /build/entrypoint.sh

CMD ["/build/entrypoint.sh"]

