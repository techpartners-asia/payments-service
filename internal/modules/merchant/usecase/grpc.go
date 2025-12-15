package usecaseMerchant

import (
	"context"

	repositoryRedis "git.techpartners.asia/gateway-services/payment-service/infrastructure/redis/repository"
	usecaseRedis "git.techpartners.asia/gateway-services/payment-service/infrastructure/redis/usecase"
	grpcMerchantRequestDTO "git.techpartners.asia/gateway-services/payment-service/internal/modules/merchant/dto/request/grpc"
	grpcMerchantResponseDTO "git.techpartners.asia/gateway-services/payment-service/internal/modules/merchant/dto/response/grpc"
	repositoryMerchant "git.techpartners.asia/gateway-services/payment-service/internal/modules/merchant/repository"
	merchantProto "git.techpartners.asia/gateway-services/payment-service/pkg/proto/merchant"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type merchantUsecase struct {
	merchantProto.UnimplementedMerchantServiceServer
	merchantRepo repositoryMerchant.MerchantRepository
	ebarimtRepo  repositoryMerchant.EbarimtRepository

	merchantRedis usecaseRedis.MerchantUsecase
}

func NewMerchantUsecase(merchantRepo repositoryMerchant.MerchantRepository, ebarimtRepo repositoryMerchant.EbarimtRepository, redisRepository repositoryRedis.RedisRepository) merchantProto.MerchantServiceServer {
	return &merchantUsecase{
		merchantRepo:  merchantRepo,
		ebarimtRepo:   ebarimtRepo,
		merchantRedis: usecaseRedis.NewMerchantUsecase(redisRepository),
	}
}

func (u *merchantUsecase) Create(ctx context.Context, req *merchantProto.CreateMerchantRequest) (*merchantProto.MerchantResponse, error) {
	merchant, err := u.merchantRepo.CreateMerchant(grpcMerchantRequestDTO.ToEntity(req))
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	ebarimtCredential, err := u.ebarimtRepo.CreateEbarimt(grpcMerchantRequestDTO.ToEbarimtEntity(req.EbarimtCredential))
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	if err := u.merchantRedis.Cache(merchant.ID, req.PaymentCredential, ebarimtCredential); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return grpcMerchantResponseDTO.ToResponse(merchant), nil
}

func (u *merchantUsecase) GetByID(ctx context.Context, req *merchantProto.MerchantIDRequest) (*merchantProto.MerchantResponse, error) {
	merchant, err := u.merchantRepo.GetMerchantByID(uint(req.Id))
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return grpcMerchantResponseDTO.ToResponse(merchant), nil
}

func (u *merchantUsecase) Update(ctx context.Context, req *merchantProto.UpdateRequest) (*merchantProto.MerchantResponse, error) {
	merchant, err := u.merchantRepo.UpdateMerchant(grpcMerchantRequestDTO.UpdateToEntity(req))
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return grpcMerchantResponseDTO.ToResponse(merchant), nil
}

func (u *merchantUsecase) Delete(ctx context.Context, req *merchantProto.MerchantIDRequest) (*merchantProto.SuccessResponse, error) {
	err := u.merchantRepo.DeleteMerchant(uint(req.Id))
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &merchantProto.SuccessResponse{Success: true}, nil
}
