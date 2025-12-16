package main

import (
	grpcCmd "git.techpartners.asia/gateway-services/payment-service/cmd/grpc"
	"git.techpartners.asia/gateway-services/payment-service/infrastructure/database"
	redisService "git.techpartners.asia/gateway-services/payment-service/infrastructure/redis"
	configPkg "git.techpartners.asia/gateway-services/payment-service/pkg/config"
)

func main() {
	configPkg.Init()
	database.Init()
	redisService.Init()
	// serverCmd.Run()
	grpcCmd.Run()
}
