package grpcMerchantRequestDTO

import (
	databaseBase "github.com/techpartners-asia/payments-service/infrastructure/database/base"
	"github.com/techpartners-asia/payments-service/infrastructure/database/entity"
	merchantProto "github.com/techpartners-asia/payments-service/pkg/proto/merchant"
)

func UpdateToEntity(req *merchantProto.UpdateRequest) *entity.MerchantEntity {
	return &entity.MerchantEntity{
		BaseEntity: databaseBase.BaseEntity{
			ID: uint(req.Id),
		},
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
