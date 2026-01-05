package grpcMerchantRequestDTO

import (
	"errors"

	"git.techpartners.asia/gateway-services/payment-service/infrastructure/database/entity"
	merchantProto "git.techpartners.asia/gateway-services/payment-service/pkg/proto/merchant"
)

func ToEntity(req *merchantProto.CreateMerchantRequest) *entity.MerchantEntity {
	return &entity.MerchantEntity{
		Name: req.Name,
		UID:  req.Uid,
	}
}

func ToEbarimtEntity(merchant *entity.MerchantEntity, req *merchantProto.MerchantEbarimtCredentialRequest) *entity.MerchantEbarimtEntity {
	return &entity.MerchantEbarimtEntity{
		MerchantID:   merchant.ID,
		Url:          req.Url,
		Tin:          req.Tin,
		PosNo:        req.PosNo,
		BranchNo:     req.BranchNo,
		DistrictCode: req.DistrictCode,
	}
}

func Validate(req *merchantProto.CreateMerchantRequest) error {
	if req.Name == "" {
		return errors.New("name is required")
	}
	if req.Uid == "" {
		return errors.New("uid is required")
	}
	if req.EbarimtCredential == nil {
		return errors.New("ebarimt credential is required")
	}
	if req.PaymentCredential == nil {
		return errors.New("payment credential is required")
	}

	return nil
}
