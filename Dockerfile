FROM iotauz/sdk:base-0.0.1

WORKDIR /build

COPY go.mod go.sum ./
RUN go get github.com/iota-uz/iota-sdk@main
RUN go mod download
COPY . .
RUN make css

RUN go build -o run_server cmd/server/main.go && go build -o seed_db cmd/seed/main.go
CMD go run cmd/migrate/main.go up && /build/seed_db && /build/run_server
