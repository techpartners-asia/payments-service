package repositoryMerchant

import "github.com/techpartners-asia/payments-service/infrastructure/database/entity"

type MerchantRepository interface {
	CreateMerchant(merchant *entity.MerchantEntity) (*entity.MerchantEntity, error)
	GetMerchantByID(id uint) (*entity.MerchantEntity, error)
	UpdateMerchant(merchant *entity.MerchantEntity) (*entity.MerchantEntity, error)
	DeleteMerchant(id uint) error
}
