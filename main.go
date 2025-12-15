package main

import (
	grpcCmd "git.techpartners.asia/gateway-services/payment-service/cmd/grpc"
	serverCmd "git.techpartners.asia/gateway-services/payment-service/cmd/server"
	"git.techpartners.asia/gateway-services/payment-service/infrastructure/database"
	redisService "git.techpartners.asia/gateway-services/payment-service/infrastructure/redis"
	configPkg "git.techpartners.asia/gateway-services/payment-service/pkg/config"
)

func main() {
	configPkg.Init()
	database.Init()
	redisService.Init()
	serverCmd.Run()
	go func() {
		grpcCmd.Run()
	}()
}
