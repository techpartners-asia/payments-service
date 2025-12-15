package repositoryImpl

import (
	"git.techpartners.asia/gateway-services/payment-service/infrastructure/database/entity"
	"gorm.io/gorm"
)

type MerchantEbarimtRepositoryImpl struct {
	db *gorm.DB
}

func NewMerchantEbarimtRepository(db *gorm.DB) *MerchantEbarimtRepositoryImpl {
	return &MerchantEbarimtRepositoryImpl{db: db}
}

func (r *MerchantEbarimtRepositoryImpl) CreateEbarimt(merchantEbarimt *entity.MerchantEbarimtEntity) (*entity.MerchantEbarimtEntity, error) {
	if err := r.db.Create(merchantEbarimt).Error; err != nil {
		return nil, err
	}
	return merchantEbarimt, nil
}

func (r *MerchantEbarimtRepositoryImpl) GetEbarimtByID(id uint) (*entity.MerchantEbarimtEntity, error) {
	var merchantEbarimt entity.MerchantEbarimtEntity
	if err := r.db.Where("id = ?", id).First(&merchantEbarimt).Error; err != nil {
		return nil, err
	}
	return &merchantEbarimt, nil
}

func (r *MerchantEbarimtRepositoryImpl) UpdateEbarimt(merchantEbarimt *entity.MerchantEbarimtEntity) (*entity.MerchantEbarimtEntity, error) {
	if err := r.db.Save(merchantEbarimt).Error; err != nil {
		return nil, err
	}
	return merchantEbarimt, nil
}

func (r *MerchantEbarimtRepositoryImpl) DeleteEbarimt(id uint) error {
	if err := r.db.Delete(&entity.MerchantEbarimtEntity{}, id).Error; err != nil {
		return err
	}
	return nil
}
