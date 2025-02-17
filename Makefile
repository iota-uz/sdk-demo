# Variables
TAILWIND_INPUT := internal/assets/css/main.css
TAILWIND_OUTPUT := internal/assets/css/main.min.css

# Install dependencies
deps:
	go get ./...

# Seed database
seed:
	go run cmd/seed/main.go

generate:
	go generate ./... && templ generate

# Apply database migrations (up)
migrate-up:
	go run cmd/migrate/main.go up

# Upgrade iota-sdk version && push
release:
	go get github.com/iota-uz/iota-sdk@main && git add . && git commit -m "Upgrade iota-sdk version" && git push

# Downgrade database migrations (down)
migrate-down:
	go run cmd/migrate/main.go down

# Run PostgreSQL
localdb:
	docker compose -f compose.dev.yml up db

clear-localdb:
	rm -rf postgres-data/

# Compile TailwindCSS (with watch)
css-watch:
	tailwindcss -c tailwind.config.js -i $(TAILWIND_INPUT) -o $(TAILWIND_OUTPUT) --minify --watch

# Compile TailwindCSS (without watch)
css:
	tailwindcss -c tailwind.config.js -i $(TAILWIND_INPUT) -o $(TAILWIND_OUTPUT) --minify

# Run linter
lint:
	golangci-lint run ./...

# Clean build artifacts
clean:
	rm -rf $(TAILWIND_OUTPUT)

# Upgrade iota-sdk version
upgrade-sdk:
	chmod +x upgrade.sh && ./upgrade.sh

# Full setup
setup: deps migrate-up css lint

.PHONY: deps localdb migrate-up migrate-down dev css-watch css lint clean setup
