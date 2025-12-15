package repositoryMerchant

import (
	"git.techpartners.asia/gateway-services/payment-service/infrastructure/database/entity"
)

type EbarimtRepository interface {
	CreateEbarimt(ebarimt *entity.MerchantEbarimtEntity) (*entity.MerchantEbarimtEntity, error)
	GetEbarimtByID(id uint) (*entity.MerchantEbarimtEntity, error)
	UpdateEbarimt(ebarimt *entity.MerchantEbarimtEntity) (*entity.MerchantEbarimtEntity, error)
	DeleteEbarimt(id uint) error
}
