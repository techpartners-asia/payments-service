package usecaseRedis

import (
	"errors"
	"fmt"

	"git.techpartners.asia/gateway-services/payment-service/infrastructure/database/entity"
	redisDTO "git.techpartners.asia/gateway-services/payment-service/infrastructure/redis/dto"
	repositoryRedis "git.techpartners.asia/gateway-services/payment-service/infrastructure/redis/repository"
	merchantProto "git.techpartners.asia/gateway-services/payment-service/pkg/proto/merchant"
	"github.com/redis/go-redis/v9"
)

const (
	MerchantCacheKey = "merchant_cache_%v"
)

type MerchantUsecase interface {
	Cache(merchantUID string, paymentCredentials *merchantProto.MerchantPaymentCredentialRequest, ebarimt *entity.MerchantEbarimtEntity) error
	Get(merchantUID string) (*redisDTO.RedisCacheDTO, error)
	Remove(merchantUID string) error
}

type merchantUsecase struct {
	redisRepository repositoryRedis.RedisRepository
}

func NewMerchantUsecase(redisRepository repositoryRedis.RedisRepository) MerchantUsecase {
	return &merchantUsecase{redisRepository: redisRepository}
}

func (u *merchantUsecase) Cache(merchantUID string, paymentCredentials *merchantProto.MerchantPaymentCredentialRequest, ebarimt *entity.MerchantEbarimtEntity) error {

	if err := u.redisRepository.Delete(fmt.Sprintf(MerchantCacheKey, merchantUID)); err != nil && err != redis.Nil {
		return err
	}

	paymentCacheDTO := redisDTO.ToRedisCacheDTO(merchantUID, paymentCredentials, ebarimt)
	if paymentCacheDTO == nil {
		return errors.New("failed to create payment cache")
	}
	return u.redisRepository.Set(fmt.Sprintf(MerchantCacheKey, merchantUID), paymentCacheDTO, 0)
}

func (u *merchantUsecase) Get(merchantUID string) (*redisDTO.RedisCacheDTO, error) {
	var cache redisDTO.RedisCacheDTO
	if err := u.redisRepository.Get(fmt.Sprintf(MerchantCacheKey, merchantUID), &cache); err != nil {
		return nil, err
	}

	return &cache, nil
}

func (u *merchantUsecase) Remove(merchantUID string) error {
	return u.redisRepository.Delete(fmt.Sprintf(MerchantCacheKey, merchantUID))
}
