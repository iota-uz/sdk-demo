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

# Downgrade database migrations (down)
migrate-down:
	go run cmd/migrate/main.go down

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

# Full setup
setup: deps localdb migrate-up css lint

.PHONY: default deps test test-watch localdb migrate-up migrate-down dev css-watch css lint clean setup
