package grpcMerchantRequestDTO

import (
	"git.techpartners.asia/gateway-services/payment-service/infrastructure/database/entity"
	merchantProto "git.techpartners.asia/gateway-services/payment-service/pkg/proto/merchant"
)

func ToEntity(req *merchantProto.CreateMerchantRequest) *entity.MerchantEntity {
	return &entity.MerchantEntity{
		Name: req.Name,
	}
}

func ToEbarimtEntity(req *merchantProto.MerchantEbarimtCredentialRequest) *entity.MerchantEbarimtEntity {
	return &entity.MerchantEbarimtEntity{
		Url:          req.Url,
		Tin:          req.Tin,
		PosNo:        req.PosNo,
		BranchNo:     req.BranchNo,
		DistrictCode: req.DistrictCode,
	}
}
