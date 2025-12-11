package repositoryMerchant

import (
	"github.com/techpartners-asia/payments-service/infrastructure/database/entity"
)

type EbarimtRepository interface {
	CreateEbarimt(ebarimt *entity.MerchantEbarimtEntity) (*entity.MerchantEbarimtEntity, error)
	GetEbarimtByID(id uint) (*entity.MerchantEbarimtEntity, error)
	UpdateEbarimt(ebarimt *entity.MerchantEbarimtEntity) (*entity.MerchantEbarimtEntity, error)
	DeleteEbarimt(id uint) error
}
