package serverCmd

import (
	httpDeliveryRoutes "git.techpartners.asia/gateway-services/payment-service/internal/delivery/http/routes"
	configPkg "git.techpartners.asia/gateway-services/payment-service/pkg/config"
	fiberPkg "git.techpartners.asia/gateway-services/payment-service/pkg/fiber"
)

func Run() {

	app := fiberPkg.NewFiber()

	httpDeliveryRoutes.Routes(app)

	if err := app.Listen(":" + configPkg.Env.App.Port); err != nil {
		panic(err)
	}
}
