# CLAUDE.md - IOTA SDK Guide

## Overview
This is a project built on [IOTA-SDK](https://github.com/iota-uz/iota-sdk).
To lookup IOTA SDK documentation, please refer `LLMS.md` file.

The project follows DDD principles. The project is divided into 3 main layers:
- Domain: Contains the business logic and domain models
- Infrastructure: Contains the database and other external services
- Application: Contains the use cases and application services

## Build/Lint/Test Commands
- After changes to Go code: `go vet ./...`
- Run all tests: `make test` or `go test -v ./...` 
- Run single test: `go test -v ./path/to/package -run TestName`
- Run specific subtest: `go test -v ./path/to/package -run TestName/SubtestName`
- JSON linting: `make build-iota-linter && make run-iota-linter`
- Apply migrations: `make migrate up`
- After changing `.templ` files DO NOT run generate commands, it is automatically done in background

## Code Style Guidelines
- DO NOT COMMENT EXECESSIVELY. Instead, write clear and concise code that is self-explanatory
- DO NOT DIRECTLY EDIT FILES under `migrations/`
- For package aliases use flat case (e.g. `import useraggregate "github.com/iota-uz/iota-sdk/domain/user"`)
- Use Go v1.23.2 and follow standard Go idioms
- File organization: group related functionality in modules/ or pkg/ directories
- Naming: use camelCase for variables, PascalCase for exported functions/types
- Testing: table-driven tests with descriptive names (TestFunctionName_Scenario)
- Error handling: use pkg/serrors for standard error types
- Type safety: use strong typing and avoid interface{} where possible
- Follow existing patterns for database operations with jmoiron/sqlx
- For UI components, follow the existing templ/htmx patterns
