FROM golang:1.23.2 AS install-stage

WORKDIR /build

RUN case "$(uname -m)" in \
      "x86_64") ARCH="x64" ;; \
      "aarch64") ARCH="arm64" ;; \
      *) echo "Unsupported architecture: $(uname -m)" && exit 1 ;; \
    esac && \
    curl -sLO https://github.com/tailwindlabs/tailwindcss/releases/download/v3.4.15/tailwindcss-linux-${ARCH} && \
    chmod +x tailwindcss-linux-${ARCH} && \
    mv tailwindcss-linux-${ARCH} /usr/local/bin/tailwindcss

RUN go install github.com/a-h/templ/cmd/templ@v0.2.793
COPY go.mod go.sum ./
RUN go get github.com/iota-uz/iota-sdk@main
RUN go mod download
COPY iota-erp .
RUN make generate && go vet ./...
RUN make css

FROM install-stage AS staging
RUN go build -o run_server cmd/server/main.go && go build -o seed_db cmd/seed/main.go
CMD go run cmd/migrate/main.go up && /build/seed_db && /build/run_server
