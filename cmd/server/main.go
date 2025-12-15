package serverCmd

import (
	httpDeliveryRoutes "github.com/techpartners-asia/payments-service/internal/delivery/http/routes"
	configPkg "github.com/techpartners-asia/payments-service/pkg/config"
	fiberPkg "github.com/techpartners-asia/payments-service/pkg/fiber"
)

func Run() {

	app := fiberPkg.NewFiber()

	httpDeliveryRoutes.Routes(app)

	if err := app.Listen(":" + configPkg.Env.App.Port); err != nil {
		panic(err)
	}
}
