package main

import (
	grpcCmd "github.com/techpartners-asia/payments-service/cmd/grpc"
	serverCmd "github.com/techpartners-asia/payments-service/cmd/server"
	"github.com/techpartners-asia/payments-service/infrastructure/database"
	redisService "github.com/techpartners-asia/payments-service/infrastructure/redis"
	configPkg "github.com/techpartners-asia/payments-service/pkg/config"
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
