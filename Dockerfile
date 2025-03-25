FROM iotauz/sdk:base-0.0.1

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN make css

RUN go build -o run_server cmd/server/main.go && go build -o seed_db cmd/seed/main.go
RUN go build -o migrate cmd/migrate/main.go
CMD /build/migrate up && /build/seed_db && /build/run_server

