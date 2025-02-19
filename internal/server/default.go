package server

import (
	"github.com/iota-uz/iota-sdk/modules/core/presentation/controllers"
	"github.com/iota-uz/iota-sdk/modules/core/presentation/templates/layouts"
	"github.com/iota-uz/iota-sdk/pkg/application"
	"github.com/iota-uz/iota-sdk/pkg/configuration"
	"github.com/iota-uz/iota-sdk/pkg/constants"
	"github.com/iota-uz/iota-sdk/pkg/middleware"
	"github.com/iota-uz/iota-sdk/pkg/server"
	internallayouts "github.com/iota-uz/sdk-demo/internal/templates/layouts"
	internalcontrollers "github.com/iota-uz/sdk-demo/modules/custom/presentation/controllers"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
)

type DefaultOptions struct {
	Logger        *logrus.Logger
	Configuration *configuration.Configuration
	Application   application.Application
	Pool          *pgxpool.Pool
}

func Default(options *DefaultOptions) (*server.HTTPServer, error) {
	app := options.Application
	app.RegisterMiddleware(
		middleware.WithLogger(options.Logger),
		middleware.Provide(constants.AppKey, app),
		middleware.Provide(constants.HeadKey, internallayouts.Head()),
		middleware.Provide(constants.LogoKey, layouts.DefaultLogo()),
		middleware.Provide(constants.PoolKey, options.Pool),
		middleware.Cors("http://localhost:3000", "ws://localhost:3000"),
		middleware.RequestParams(),
		middleware.LogRequests(),
	)
	app.RegisterControllers(internalcontrollers.NewLoginController(app))
	serverInstance := server.NewHTTPServer(
		app,
		controllers.NotFound(options.Application),
		controllers.MethodNotAllowed(),
	)
	return serverInstance, nil
}
