# IOTA-SDK Demo

This repository demonstrates how to use and customize the **IOTA-SDK**, an open-source modular ERP system, for your
specific needs. Explore examples, learn integration techniques, and understand the framework’s extensibility.

---

## Table of contents
- [Quick-start](#quick-start)
- [Modules](#modules)
- [What's inside](#what-s-inside)

---

## 🚀 Quick Start

1. **Clone the Repository**:
   ```bash
   git clone https://github.com/iota-uz/sdk-demo.git
   cd iota-sdk-demo
   ```

2. **Install Dependencies**:
   ```bash
   make deps
   ```

3. **Run the application**:
   ```bash
   docker compose -f docker-compose.dev.yml up
   make migrate-up
   ```

4. **View the Demo**:
   Navigate to `http://localhost:3200` in your browser.

---

## Modules

IOTA-SDK comes with a few built-in modules: 

- **`core`** - contains authentication & authorization, file uploads, and user management
- **`finance`** - contains expenses and payments
- **`warehouse`** - contains orders, products, and inventory management

Registering a module is quite simple:
1. **Create your application**

```go
import (
  "github.com/iota-uz/iota-sdk/pkg/application"
  "github.com/iota-uz/iota-sdk/pkg/configuration"
  "github.com/iota-uz/iota-sdk/pkg/eventbus"
  "github.com/iota-uz/iota-sdk/pkg/modules"
  "github.com/iota-uz/iota-sdk/pkg/modules/core"
  "github.com/iota-uz/iota-sdk/pkg/modules/finance"
  "github.com/jackc/pgx/v5/pgxpool"
  "context"
)
func main() {
  conf := configuration.Use()
  pool, err := pgxpool.New(context.TODO(), conf.DBOpts)
  app := application.New(pool, eventbus.NewEventPublisher())
}
```
2. **Register modules**
```go
// code from step 1
if err := modules.Load(app, core.NewModule(), finance.NewModule()); err != nil {
  log.Fatalf("failed to load modules: %v", err)
}
```

If you wish to register all built-in modules at once, use `modules.BuiltInModules`

```go
if err := modules.Load(app, modules.BuiltInModules...); err != nil {
  log.Fatalf("failed to load modules: %v", err)
}
```

---

## 🧩 What’s Inside?

- **Usage Examples**: Illustrates modularity, templating, and API integration.
- **Customization Samples**: Learn to extend IOTA-SDK with custom modules and workflows.
- **Frontend Integration**: Explore HTMX and Alpine.js for reactive components.
