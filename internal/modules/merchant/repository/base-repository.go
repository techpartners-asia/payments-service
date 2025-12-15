package repositoryMerchant

import "git.techpartners.asia/gateway-services/payment-service/infrastructure/database/entity"

type MerchantRepository interface {
	CreateMerchant(merchant *entity.MerchantEntity) (*entity.MerchantEntity, error)
	GetMerchantByID(id uint) (*entity.MerchantEntity, error)
	UpdateMerchant(merchant *entity.MerchantEntity) (*entity.MerchantEntity, error)
	DeleteMerchant(id uint) error
}
