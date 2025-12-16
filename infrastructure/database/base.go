package database

import (
	"fmt"
	"log"
	"sync"

	"git.techpartners.asia/gateway-services/payment-service/infrastructure/database/entity"
	configPkg "git.techpartners.asia/gateway-services/payment-service/pkg/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB   *gorm.DB
	once sync.Once
)

func Init() {
	fmt.Println(configPkg.Env.DB)
	once.Do(func() {
		db, err := gorm.Open(postgres.Open(
			fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable TimeZone=%s",
				configPkg.Env.DB.Host,
				configPkg.Env.DB.Port,
				configPkg.Env.DB.User,
				configPkg.Env.DB.Name,
				configPkg.Env.DB.Password,
				configPkg.Env.DB.Timezone,
			),
		),
			&gorm.Config{
				PrepareStmt:                              true,
				SkipDefaultTransaction:                   true,
				DisableForeignKeyConstraintWhenMigrating: true,
			},
		)
		if err != nil {
			log.Fatalf("Failed to connect to database: %v", err)
		}

		if err := db.AutoMigrate(
			&entity.PaymentEntity{},
			&entity.MerchantEntity{},
			&entity.MerchantEbarimtEntity{},
			&entity.EbarimtEntity{},
			&entity.EbarimtReceiptEntity{},
			&entity.EbarimtReceiptItemEntity{},
		); err != nil {
			log.Fatalf("Failed to migrate database: %v", err)
		}

		DB = db
	})
}
