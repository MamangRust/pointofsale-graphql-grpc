package graphql

import (
	"github.com/MamangRust/pointofsale-graphql-grpc/internal/model"
	"github.com/MamangRust/pointofsale-graphql-grpc/internal/pb"
)

type merchantGraphqlMapper struct {
}

func NewMerchantGraphqlMapper() *merchantGraphqlMapper {
	return &merchantGraphqlMapper{}
}

func (m *merchantGraphqlMapper) ToGraphqlResponseMerchant(resp *pb.ApiResponseMerchant) *model.APIResponseMerchant {
	if resp == nil || resp.Data == nil {
		return nil
	}

	return &model.APIResponseMerchant{
		Status:  resp.Status,
		Message: resp.Message,
		Data:    m.mapMerchantResponse(resp.Data),
	}
}

func (m *merchantGraphqlMapper) ToGraphqlResponseMerchantDeleteAt(resp *pb.ApiResponseMerchantDeleteAt) *model.APIResponseMerchantDeleteAt {
	if resp == nil || resp.Data == nil {
		return nil
	}

	return &model.APIResponseMerchantDeleteAt{
		Status:  resp.Status,
		Message: resp.Message,
		Data:    m.mapMerchantResponseDeleteAt(resp.Data),
	}
}

func (m *merchantGraphqlMapper) ToGraphqlResponseMerchantAll(resp *pb.ApiResponseMerchantAll) *model.APIResponseMerchantAll {
	if resp == nil {
		return nil
	}

	return &model.APIResponseMerchantAll{
		Status:  resp.Status,
		Message: resp.Message,
	}
}

func (m *merchantGraphqlMapper) ToGraphqlResponseMerchantDelete(resp *pb.ApiResponseMerchantDelete) *model.APIResponseMerchantDelete {
	if resp == nil {
		return nil
	}

	return &model.APIResponseMerchantDelete{
		Status:  resp.Status,
		Message: resp.Message,
	}
}

func (m *merchantGraphqlMapper) ToGraphqlResponsesMerchant(resp *pb.ApiResponsesMerchant) *model.APIResponsesMerchant {
	if resp == nil {
		return nil
	}

	return &model.APIResponsesMerchant{
		Status:  resp.Status,
		Message: resp.Message,
		Data:    m.mapMerchantResponses(resp.Data),
	}
}

func (m *merchantGraphqlMapper) ToGraphqlResponsePaginationMerchant(resp *pb.ApiResponsePaginationMerchant) *model.APIResponsePaginationMerchant {
	if resp == nil {
		return nil
	}

	return &model.APIResponsePaginationMerchant{
		Status:     resp.Status,
		Message:    resp.Message,
		Data:       m.mapMerchantResponses(resp.Data),
		Pagination: mapPaginationMeta(resp.Pagination),
	}
}

func (m *merchantGraphqlMapper) ToGraphqlResponsePaginationMerchantDeleteAt(resp *pb.ApiResponsePaginationMerchantDeleteAt) *model.APIResponsePaginationMerchantDeleteAt {
	if resp == nil {
		return nil
	}

	return &model.APIResponsePaginationMerchantDeleteAt{
		Status:     resp.Status,
		Message:    resp.Message,
		Data:       m.mapMerchantResponsesDeleteAt(resp.Data),
		Pagination: mapPaginationMeta(resp.Pagination),
	}
}

func (m *merchantGraphqlMapper) mapMerchantResponse(merchant *pb.MerchantResponse) *model.MerchantResponse {
	if merchant == nil {
		return nil
	}

	return &model.MerchantResponse{
		ID:           int32(merchant.Id),
		UserID:       int32(merchant.UserId),
		Name:         merchant.Name,
		Description:  &merchant.Description,
		Address:      &merchant.Address,
		ContactEmail: &merchant.ContactEmail,
		ContactPhone: &merchant.ContactPhone,
		Status:       &merchant.Status,
		CreatedAt:    &merchant.CreatedAt,
		UpdatedAt:    &merchant.UpdatedAt,
	}
}

func (m *merchantGraphqlMapper) mapMerchantResponses(merchants []*pb.MerchantResponse) []*model.MerchantResponse {
	res := make([]*model.MerchantResponse, 0, len(merchants))
	for _, merchant := range merchants {
		res = append(res, m.mapMerchantResponse(merchant))
	}
	return res
}

func (m *merchantGraphqlMapper) mapMerchantResponseDeleteAt(merchant *pb.MerchantResponseDeleteAt) *model.MerchantResponseDeleteAt {
	if merchant == nil {
		return nil
	}

	var deletedAt string

	if merchant.DeletedAt != nil {
		deletedAt = merchant.DeletedAt.Value
	}

	return &model.MerchantResponseDeleteAt{
		ID:           int32(merchant.Id),
		UserID:       int32(merchant.UserId),
		Name:         merchant.Name,
		Description:  &merchant.Description,
		Address:      &merchant.Address,
		ContactEmail: &merchant.ContactEmail,
		ContactPhone: &merchant.ContactPhone,
		Status:       &merchant.Status,
		CreatedAt:    &merchant.CreatedAt,
		UpdatedAt:    &merchant.UpdatedAt,
		DeletedAt:    &deletedAt,
	}
}

func (m *merchantGraphqlMapper) mapMerchantResponsesDeleteAt(merchants []*pb.MerchantResponseDeleteAt) []*model.MerchantResponseDeleteAt {
	res := make([]*model.MerchantResponseDeleteAt, 0, len(merchants))
	for _, merchant := range merchants {
		res = append(res, m.mapMerchantResponseDeleteAt(merchant))
	}
	return res
}
