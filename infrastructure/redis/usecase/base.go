package usecaseRedis

import (
	"errors"
	"fmt"

	"github.com/techpartners-asia/payments-service/infrastructure/database/entity"
	redisDTO "github.com/techpartners-asia/payments-service/infrastructure/redis/dto"
	repositoryRedis "github.com/techpartners-asia/payments-service/infrastructure/redis/repository"
	merchantProto "github.com/techpartners-asia/payments-service/pkg/proto/merchant"
)

const (
	MerchantCacheKey = "merchant_cache_%d"
)

type MerchantUsecase interface {
	Cache(merchantID uint, paymentCredentials *merchantProto.MerchantPaymentCredentialRequest, ebarimt *entity.MerchantEbarimtEntity) error
	Get(merchantID uint) (*redisDTO.RedisCacheDTO, error)
	Remove(merchantID uint) error
}

type merchantUsecase struct {
	redisRepository repositoryRedis.RedisRepository
}

func NewMerchantUsecase(redisRepository repositoryRedis.RedisRepository) MerchantUsecase {
	return &merchantUsecase{redisRepository: redisRepository}
}

func (u *merchantUsecase) Cache(merchantID uint, paymentCredentials *merchantProto.MerchantPaymentCredentialRequest, ebarimt *entity.MerchantEbarimtEntity) error {
	paymentCacheDTO := redisDTO.ToRedisCacheDTO(merchantID, paymentCredentials, ebarimt)
	if paymentCacheDTO == nil {
		return errors.New("failed to create payment cache")
	}
	return u.redisRepository.Set(fmt.Sprintf(MerchantCacheKey, merchantID), paymentCacheDTO, 0)
}

func (u *merchantUsecase) Get(merchantID uint) (*redisDTO.RedisCacheDTO, error) {
	var cache redisDTO.RedisCacheDTO
	if err := u.redisRepository.Get(fmt.Sprintf(MerchantCacheKey, merchantID), &cache); err != nil {
		return nil, err
	}

	return &cache, nil
}

func (u *merchantUsecase) Remove(merchantID uint) error {
	return u.redisRepository.Delete(fmt.Sprintf(MerchantCacheKey, merchantID))
}
