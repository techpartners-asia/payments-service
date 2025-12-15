package repositoryImpl

import (
	"git.techpartners.asia/gateway-services/payment-service/infrastructure/database/entity"
	"gorm.io/gorm"
)

type MerchantRepositoryImpl struct {
	db *gorm.DB
}

func NewMerchantRepository(db *gorm.DB) *MerchantRepositoryImpl {
	return &MerchantRepositoryImpl{db: db}
}

func (r *MerchantRepositoryImpl) CreateMerchant(merchant *entity.MerchantEntity) (*entity.MerchantEntity, error) {
	if err := r.db.Create(merchant).Error; err != nil {
		return nil, err
	}
	return merchant, nil
}

func (r *MerchantRepositoryImpl) GetMerchantByID(id uint) (*entity.MerchantEntity, error) {
	var merchant entity.MerchantEntity
	if err := r.db.Where("id = ?", id).First(&merchant).Error; err != nil {
		return nil, err
	}

	return &merchant, nil
}

func (r *MerchantRepositoryImpl) UpdateMerchant(merchant *entity.MerchantEntity) (*entity.MerchantEntity, error) {
	if err := r.db.Save(merchant).Error; err != nil {
		return nil, err
	}
	return merchant, nil
}

func (r *MerchantRepositoryImpl) DeleteMerchant(id uint) error {
	if err := r.db.Delete(&entity.MerchantEntity{}, id).Error; err != nil {
		return err
	}
	return nil
}
