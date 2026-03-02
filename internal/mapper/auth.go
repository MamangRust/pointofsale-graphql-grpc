package graphql

import (
	"github.com/MamangRust/pointofsale-graphql-grpc/internal/model"
	"github.com/MamangRust/pointofsale-graphql-grpc/internal/pb"
)

type authResponseMapper struct {
}

func NewAuthResponseMapper() *authResponseMapper {
	return &authResponseMapper{}
}

func (s *authResponseMapper) ToGraphqlResponseLogin(res *pb.ApiResponseLogin) *model.APIResponseLogin {
	return &model.APIResponseLogin{
		Status:  res.Status,
		Message: res.Message,
		Data:    s.mapResponseToken(res.Data),
	}
}

func (s *authResponseMapper) ToGraphqlResponseRegister(res *pb.ApiResponseRegister) *model.APIResponseRegister {
	return &model.APIResponseRegister{
		Status:  res.Status,
		Message: res.Message,
		Data:    s.mapResponseUser(res.Data),
	}
}

func (s *authResponseMapper) ToGraphqlResponseRefreshToken(res *pb.ApiResponseRefreshToken) *model.APIResponseRefreshToken {
	return &model.APIResponseRefreshToken{
		Status:  res.Status,
		Message: res.Message,
		Data:    s.mapResponseToken(res.Data),
	}
}

func (s *authResponseMapper) ToGraphqlResponseGetMe(res *pb.ApiResponseGetMe) *model.APIResponseGetMe {
	return &model.APIResponseGetMe{
		Status:  res.Status,
		Message: res.Message,
		Data:    s.mapResponseUser(res.Data),
	}
}

func (s *authResponseMapper) mapResponseUser(res *pb.UserResponse) *model.UserResponse {
	return &model.UserResponse{
		ID:        res.Id,
		Firstname: res.Firstname,
		Lastname:  res.Lastname,
		Email:     res.Email,
		CreatedAt: res.CreatedAt,
		UpdatedAt: res.UpdatedAt,
	}
}

func (s *authResponseMapper) mapResponseToken(res *pb.TokenResponse) *model.TokenResponse {
	return &model.TokenResponse{
		AccessToken:  res.AccessToken,
		RefreshToken: res.RefreshToken,
	}
}
