package database

import (
	"fmt"
	"log"
	"sync"

	configPkg "github.com/techpartners-asia/payments-service/pkg/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB   *gorm.DB
	once sync.Once
)

func Init() {
	once.Do(func() {
		db, err := gorm.Open(postgres.Open(
			fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
				configPkg.Env.DB.Host,
				configPkg.Env.DB.Port,
				configPkg.Env.DB.User,
				configPkg.Env.DB.Password,
				configPkg.Env.DB.Name,
			),
		),
			&gorm.Config{},
		)
		if err != nil {
			log.Fatalf("Failed to connect to database: %v", err)
		}

		DB = db
	})
}
