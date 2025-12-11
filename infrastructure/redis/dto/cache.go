package redisDTO

import (
	"github.com/techpartners-asia/payments-service/infrastructure/database/entity"
	sharedDTO "github.com/techpartners-asia/payments-service/infrastructure/shared"
	merchantProto "github.com/techpartners-asia/payments-service/pkg/proto/merchant"
)

type (
	RedisCacheDTO struct {
		ID      uint                             `json:"id"`
		Configs sharedDTO.SharedPaymentConfigDTO `json:"configs"`
		Ebarimt *RedisCacheEbarimtDTO            `json:"ebarimt"`
	}

	RedisCacheEbarimtDTO struct {
		Url          string `gorm:"url" json:"url"`
		Tin          string `gorm:"tin" json:"tin"`
		PosNo        string `gorm:"pos_no" json:"pos_no"`
		BranchNo     string `gorm:"branch_no" json:"branch_no"`
		DistrictCode string `gorm:"district_code" json:"district_code"`
	}
)

func ToRedisCacheDTO(merchantID uint, paymentCredential *merchantProto.MerchantPaymentCredentialRequest, ebarimt *entity.MerchantEbarimtEntity) *RedisCacheDTO {
	return &RedisCacheDTO{
		ID:      merchantID,
		Configs: ToRedisCachePaymentDTO(paymentCredential),
		Ebarimt: ToRedisCacheEbarimtDTO(ebarimt),
	}
}

func ToRedisCacheEbarimtDTO(ebarimt *entity.MerchantEbarimtEntity) *RedisCacheEbarimtDTO {
	return &RedisCacheEbarimtDTO{
		Url:          ebarimt.Url,
		Tin:          ebarimt.Tin,
		PosNo:        ebarimt.PosNo,
		BranchNo:     ebarimt.BranchNo,
		DistrictCode: ebarimt.DistrictCode,
	}
}

func ToRedisCachePaymentDTO(req *merchantProto.MerchantPaymentCredentialRequest) sharedDTO.SharedPaymentConfigDTO {
	var configs sharedDTO.SharedPaymentConfigDTO
	if req.Qpay != nil {
		configs.Qpay = sharedDTO.QpayAdapterDTO{
			Username:    req.Qpay.Username,
			Password:    req.Qpay.Password,
			Endpoint:    req.Qpay.Endpoint,
			Callback:    req.Qpay.Callback,
			InvoiceCode: req.Qpay.InvoiceCode,
			MerchantID:  req.Qpay.MerchantId,
		}
	}
	if req.Tokipay != nil {
		configs.Tokiay = sharedDTO.TokipayAdapterDTO{
			Endpoint:      req.Tokipay.Endpoint,
			APIKey:        req.Tokipay.ApiKey,
			IMAPIKey:      req.Tokipay.ImApiKey,
			Authorization: req.Tokipay.Authorization,
			MerchantID:    req.Tokipay.MerchantId,
			SuccessURL:    req.Tokipay.SuccessUrl,
			FailureURL:    req.Tokipay.FailureUrl,
			AppSchemaIOS:  req.Tokipay.AppSchemaIos,
		}
	}
	if req.Monpay != nil {
		configs.Monpay = sharedDTO.MonpayAdapterDTO{
			Endpoint:  req.Monpay.Endpoint,
			Username:  req.Monpay.Username,
			AccountID: req.Monpay.AccountId,
			Callback:  req.Monpay.Callback,
		}
	}
	if req.Golomt != nil {
		configs.Golomt = sharedDTO.GolomtAdapterDTO{
			BaseURL:     req.Golomt.BaseUrl,
			Secret:      req.Golomt.Secret,
			BearerToken: req.Golomt.BearerToken,
		}
	}
	if req.Socialpay != nil {
		configs.Socialpay = sharedDTO.SocialPayAdapterDTO{
			Terminal: req.Socialpay.Terminal,
			Secret:   req.Socialpay.Secret,
			Endpoint: req.Socialpay.Endpoint,
		}
	}
	if req.Storepay != nil {
		configs.Storepay = sharedDTO.StorePayAdapterDTO{
			AppUserName: req.Storepay.AppUserName,
			AppPassword: req.Storepay.AppPassword,
			Username:    req.Storepay.Username,
			Password:    req.Storepay.Password,
			AuthUrl:     req.Storepay.AuthUrl,
			BaseUrl:     req.Storepay.BaseUrl,
			StoreId:     req.Storepay.StoreId,
			CallbackUrl: req.Storepay.CallbackUrl,
		}
	}
	if req.Pocket != nil {
		configs.Pocket = sharedDTO.PocketAdapterDTO{
			Merchant:      req.Pocket.Merchant,
			ClientID:      req.Pocket.ClientId,
			ClientSecret:  req.Pocket.ClientSecret,
			Environment:   req.Pocket.Environment,
			TerminalIDRaw: req.Pocket.TerminalIdRaw,
		}
	}
	if req.Simple != nil {
		configs.Simple = sharedDTO.SimpleAdapterDTO{
			UserName:    req.Simple.UserName,
			Password:    req.Simple.Password,
			BaseUrl:     req.Simple.BaseUrl,
			CallbackUrl: req.Simple.CallbackUrl,
		}
	}
	if req.Balc != nil {
		configs.Balc = sharedDTO.BalcAdapterDTO{
			Endpoint: req.Balc.Endpoint,
			Token:    req.Balc.Token,
		}
	}
	return configs
}
