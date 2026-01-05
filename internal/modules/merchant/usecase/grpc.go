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
	"gorm.io/gorm"
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

func (u *merchantUsecase) Create(ctx context.Context, req *merchantProto.CreateMerchantRequest) (*merchantProto.SuccessResponse, error) {

	if err := grpcMerchantRequestDTO.Validate(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	// TODO : Check if merchant already exists
	checkMerchant, err := u.merchantRepo.GetMerchantByUID(req.Uid)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, status.Error(codes.Internal, err.Error())
	}

	if checkMerchant != nil {
		return nil, status.Error(codes.AlreadyExists, "merchant already exists")
	}

	merchant, err := u.merchantRepo.CreateMerchant(grpcMerchantRequestDTO.ToEntity(req))
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	ebarimtCredential, err := u.ebarimtRepo.CreateEbarimt(grpcMerchantRequestDTO.ToEbarimtEntity(merchant, req.EbarimtCredential))
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	if err := u.merchantRedis.Cache(merchant.UID, req.PaymentCredential, ebarimtCredential); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &merchantProto.SuccessResponse{Success: true}, nil
}

func (u *merchantUsecase) GetByID(ctx context.Context, req *merchantProto.MerchantIDRequest) (*merchantProto.MerchantResponse, error) {
	merchant, err := u.merchantRepo.GetMerchantByUID(req.Id)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return grpcMerchantResponseDTO.ToResponse(merchant), nil
}

func (u *merchantUsecase) Update(ctx context.Context, req *merchantProto.UpdateRequest) (*merchantProto.SuccessResponse, error) {
	merchant, err := u.merchantRepo.UpdateMerchant(grpcMerchantRequestDTO.UpdateToEntity(req))
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	ebarimtCredential, err := u.ebarimtRepo.UpdateEbarimt(grpcMerchantRequestDTO.UpdateToEbarimtEntity(req))
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	if err := u.merchantRedis.Cache(merchant.UID, req.PaymentCredential, ebarimtCredential); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &merchantProto.SuccessResponse{Success: true}, nil
}

func (u *merchantUsecase) Delete(ctx context.Context, req *merchantProto.MerchantIDRequest) (*merchantProto.SuccessResponse, error) {
	err := u.merchantRepo.DeleteMerchantByUID(req.Id)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &merchantProto.SuccessResponse{Success: true}, nil
}

func (u *merchantUsecase) Save(ctx context.Context, req *merchantProto.CreateMerchantRequest) (*merchantProto.SuccessResponse, error) {
	check, err := u.merchantRepo.GetMerchantByUID(req.Uid)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, status.Error(codes.Internal, err.Error())
	}
	if check != nil {
		merchant, err := u.merchantRepo.UpdateMerchant(grpcMerchantRequestDTO.ToEntity(req))
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}

		ebarimtCredential, err := u.ebarimtRepo.UpdateEbarimt(grpcMerchantRequestDTO.ToEbarimtEntity(merchant, req.EbarimtCredential))
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}

		if err := u.merchantRedis.Cache(merchant.UID, req.PaymentCredential, ebarimtCredential); err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}

	} else {
		merchant, err := u.merchantRepo.CreateMerchant(grpcMerchantRequestDTO.ToEntity(req))
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}

		ebarimtCredential, err := u.ebarimtRepo.CreateEbarimt(grpcMerchantRequestDTO.ToEbarimtEntity(merchant, req.EbarimtCredential))
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}

		if err := u.merchantRedis.Cache(merchant.UID, req.PaymentCredential, ebarimtCredential); err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
	}
	return &merchantProto.SuccessResponse{Success: true}, nil
}
