package graphql

import (
	"github.com/MamangRust/pointofsale-graphql-grpc/internal/model"
	"github.com/MamangRust/pointofsale-graphql-grpc/internal/pb"
)

type userResponseMapper struct {
}

func NewUserResponseMapper() *userResponseMapper {
	return &userResponseMapper{}
}

func (u *userResponseMapper) ToGraphqlResponseUserDelete(res *pb.ApiResponseUserDelete) *model.APIResponseUserDelete {
	return &model.APIResponseUserDelete{
		Status:  res.Status,
		Message: res.Message,
	}
}

func (u *userResponseMapper) ToGraphqlResponseUserAll(res *pb.ApiResponseUserAll) *model.APIResponseUserAll {
	return &model.APIResponseUserAll{
		Status:  res.Status,
		Message: res.Message,
	}
}

func (u *userResponseMapper) ToGraphqlResponseUser(res *pb.ApiResponseUser) *model.APIResponseUserResponse {
	return &model.APIResponseUserResponse{
		Status:  res.Status,
		Message: res.Message,
		Data:    u.mapUserResponse(res.Data),
	}
}

func (u *userResponseMapper) ToGraphqlResponseUserDeleteAt(res *pb.ApiResponseUserDeleteAt) *model.APIResponseUserResponseDeleteAt {
	return &model.APIResponseUserResponseDeleteAt{
		Status:  res.Status,
		Message: res.Message,
		Data:    u.mapUserResponseDeleteAt(res.Data),
	}
}

func (u *userResponseMapper) ToGraphqlResponseUsers(res *pb.ApiResponsesUser) *model.APIResponsesUser {
	return &model.APIResponsesUser{
		Status:  res.Status,
		Message: res.Message,
		Data:    u.mapUserResponses(res.Data),
	}
}

func (u *userResponseMapper) ToGraphqlResponsePaginationUser(res *pb.ApiResponsePaginationUser) *model.APIResponsePaginationUser {
	return &model.APIResponsePaginationUser{
		Status:     res.Status,
		Message:    res.Message,
		Data:       u.mapUserResponses(res.Data),
		Pagination: mapPaginationMeta(res.Pagination),
	}
}

func (u *userResponseMapper) ToGraphqlResponsePaginationUserDeleteAt(res *pb.ApiResponsePaginationUserDeleteAt) *model.APIResponsePaginationUserDeleteAt {
	return &model.APIResponsePaginationUserDeleteAt{
		Status:     res.Status,
		Message:    res.Message,
		Data:       u.mapUserResponsesDeleteAt(res.Data),
		Pagination: mapPaginationMeta(res.Pagination),
	}
}

func (u *userResponseMapper) mapUserResponse(user *pb.UserResponse) *model.UserResponse {
	return &model.UserResponse{
		ID:        int32(user.Id),
		Firstname: user.Firstname,
		Lastname:  user.Lastname,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func (u *userResponseMapper) mapUserResponses(users []*pb.UserResponse) []*model.UserResponse {
	var responses []*model.UserResponse

	for _, user := range users {
		responses = append(responses, u.mapUserResponse(user))
	}

	return responses
}

func (u *userResponseMapper) mapUserResponseDeleteAt(user *pb.UserResponseDeleteAt) *model.UserResponseDeleteAt {
	var deletedAt string

	if user.DeletedAt != nil {
		deletedAt = user.DeletedAt.Value
	}

	return &model.UserResponseDeleteAt{
		ID:        int32(user.Id),
		Firstname: user.Firstname,
		Lastname:  user.Lastname,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		DeletedAt: &deletedAt,
	}
}

func (u *userResponseMapper) mapUserResponsesDeleteAt(users []*pb.UserResponseDeleteAt) []*model.UserResponseDeleteAt {
	var responses []*model.UserResponseDeleteAt

	for _, user := range users {
		responses = append(responses, u.mapUserResponseDeleteAt(user))
	}

	return responses
}

func mapPaginationMeta(s *pb.PaginationMeta) *model.PaginationMeta {
	return &model.PaginationMeta{
		CurrentPage:  int32(s.CurrentPage),
		PageSize:     int32(s.PageSize),
		TotalRecords: int32(s.TotalRecords),
		TotalPages:   int32(s.TotalPages),
	}
}
