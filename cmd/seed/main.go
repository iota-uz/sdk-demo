package main

import (
	"context"
	"runtime/debug"
	"time"

	"github.com/iota-uz/iota-sdk/modules"
	"github.com/iota-uz/iota-sdk/modules/core/domain/aggregates/user"
	"github.com/iota-uz/iota-sdk/modules/core/domain/value_objects/internet"
	coreseed "github.com/iota-uz/iota-sdk/modules/core/seed"
	"github.com/iota-uz/iota-sdk/pkg/application"
	"github.com/iota-uz/iota-sdk/pkg/composables"
	"github.com/iota-uz/iota-sdk/pkg/configuration"
	"github.com/iota-uz/iota-sdk/pkg/eventbus"

	"github.com/jackc/pgx/v5/pgxpool"
)

func panicWithStack(err error) {
	errorWithStack := string(debug.Stack()) + "\n\nError: " + err.Error()
	panic(errorWithStack)
}

func main() {
	conf := configuration.Use()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	pool, err := pgxpool.New(ctx, conf.Database.Opts)
	if err != nil {
		panic(err)
	}
	app := application.New(pool, eventbus.NewEventPublisher(conf.Logger()))
	if err := modules.Load(app); err != nil {
		panic(err)
	}
	app.RegisterNavItems(modules.NavLinks...)
	tx, err := pool.Begin(ctx)
	if err != nil {
		panic(err)
	}

	if err := modules.Load(app); err != nil {
		panicWithStack(err)
	}

	seeder := application.NewSeeder()
	email, err := internet.NewEmail("test@gmail.com")
	if err != nil {
		panicWithStack(err)
	}
	usr, err := user.New("Test", "User", email, user.UILanguageEN).SetPassword("TestPass123!")
	if err != nil {
		panicWithStack(err)
	}
	seeder.Register(
		coreseed.CreatePermissions,
		coreseed.UserSeedFunc(usr),
	)
	if err := seeder.Seed(composables.WithTx(ctx, tx), app); err != nil {
		panicWithStack(err)
	}
	if err := tx.Commit(ctx); err != nil {
		panicWithStack(err)
	}
}
