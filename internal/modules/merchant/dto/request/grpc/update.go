package grpcMerchantRequestDTO

import (
	"git.techpartners.asia/gateway-services/payment-service/infrastructure/database/entity"
	merchantProto "git.techpartners.asia/gateway-services/payment-service/pkg/proto/merchant"
)

func UpdateToEntity(req *merchantProto.UpdateRequest) *entity.MerchantEntity {
	return &entity.MerchantEntity{
		UID:  req.Uid,
		Name: req.Name,
	}
}

func UpdateToEbarimtEntity(req *merchantProto.UpdateRequest) *entity.MerchantEbarimtEntity {
	if req.EbarimtCredential == nil {
		return nil
	}

	return &entity.MerchantEbarimtEntity{
		Url:          req.EbarimtCredential.Url,
		Tin:          req.EbarimtCredential.Tin,
		PosNo:        req.EbarimtCredential.PosNo,
		BranchNo:     req.EbarimtCredential.BranchNo,
		DistrictCode: req.EbarimtCredential.DistrictCode,
	}
}
