package graphql

import (
	"github.com/MamangRust/pointofsale-graphql-grpc/internal/model"
	"github.com/MamangRust/pointofsale-graphql-grpc/internal/pb"
)

type categoryGraphqlMapper struct {
}

func NewCategoryGraphqlMapper() *categoryGraphqlMapper {
	return &categoryGraphqlMapper{}
}

func (c *categoryGraphqlMapper) ToGraphqlResponseCategory(res *pb.ApiResponseCategory) *model.APIResponseCategory {
	return &model.APIResponseCategory{
		Status:  res.Status,
		Message: res.Message,
		Data:    c.mapResponseCategory(res.Data),
	}
}

func (c *categoryGraphqlMapper) ToGraphqlResponseCategoryDeleteAt(res *pb.ApiResponseCategoryDeleteAt) *model.APIResponseCategoryDeleteAt {
	return &model.APIResponseCategoryDeleteAt{
		Status:  res.Status,
		Message: res.Message,
		Data:    c.mapResponseCategoryDeleteAt(res.Data),
	}
}

func (c *categoryGraphqlMapper) ToGraphqlResponsesCategory(res *pb.ApiResponsesCategory) *model.APIResponsesCategory {
	return &model.APIResponsesCategory{
		Status:  res.Status,
		Message: res.Message,
		Data:    c.mapResponsesCategory(res.Data),
	}
}

func (c *categoryGraphqlMapper) ToGraphqlResponseCategoryDelete(res *pb.ApiResponseCategoryDelete) *model.APIResponseCategoryDelete {
	return &model.APIResponseCategoryDelete{
		Status:  res.Status,
		Message: res.Message,
	}
}

func (c *categoryGraphqlMapper) ToGraphqlResponseCategoryAll(res *pb.ApiResponseCategoryAll) *model.APIResponseCategoryAll {
	return &model.APIResponseCategoryAll{
		Status:  res.Status,
		Message: res.Message,
	}
}

func (c *categoryGraphqlMapper) ToGraphqlResponsePaginationCategory(res *pb.ApiResponsePaginationCategory) *model.APIResponsePaginationCategory {
	return &model.APIResponsePaginationCategory{
		Status:     res.Status,
		Message:    res.Message,
		Data:       c.mapResponsesCategory(res.Data),
		Pagination: mapPaginationMeta(res.Pagination),
	}
}

func (c *categoryGraphqlMapper) ToGraphqlResponsePaginationCategoryDeleteAt(res *pb.ApiResponsePaginationCategoryDeleteAt) *model.APIResponsePaginationCategoryDeleteAt {
	return &model.APIResponsePaginationCategoryDeleteAt{
		Status:     res.Status,
		Message:    res.Message,
		Data:       c.mapResponsesCategoryDeleteAt(res.Data),
		Pagination: mapPaginationMeta(res.Pagination),
	}
}

func (c *categoryGraphqlMapper) ToGraphqlResponseCategoryMonthlyPrice(res *pb.ApiResponseCategoryMonthPrice) *model.APIResponseCategoryMonthPrice {
	return &model.APIResponseCategoryMonthPrice{
		Status:  res.Status,
		Message: res.Message,
		Data:    c.mapResponsesCategoryMonthlyPrice(res.Data),
	}
}

func (c *categoryGraphqlMapper) ToGraphqlResponseCategoryYearlyPrice(res *pb.ApiResponseCategoryYearPrice) *model.APIResponseCategoryYearPrice {
	return &model.APIResponseCategoryYearPrice{
		Status:  res.Status,
		Message: res.Message,
		Data:    c.mapResponsesCategoryYearlyPrice(res.Data),
	}
}

func (c *categoryGraphqlMapper) ToGraphqlResponseCategoryMonthlyTotalPrice(res *pb.ApiResponseCategoryMonthlyTotalPrice) *model.APIResponseCategoryMonthlyTotalPrice {
	return &model.APIResponseCategoryMonthlyTotalPrice{
		Status:  res.Status,
		Message: res.Message,
		Data:    c.mapResponsesCategoryMonthlyTotalPrice(res.Data),
	}
}

func (c *categoryGraphqlMapper) ToGraphqlResponseCategoryYearlyTotalPrice(res *pb.ApiResponseCategoryYearlyTotalPrice) *model.APIResponseCategoryYearlyTotalPrice {
	return &model.APIResponseCategoryYearlyTotalPrice{
		Status:  res.Status,
		Message: res.Message,
		Data:    c.mapResponsesCategoryYearlyTotalPrice(res.Data),
	}
}

func (c *categoryGraphqlMapper) mapResponseCategory(data *pb.CategoryResponse) *model.CategoryResponse {
	if data == nil {
		return nil
	}
	return &model.CategoryResponse{
		ID:            int32(data.Id),
		Name:          data.Name,
		Description:   &data.Description,
		SlugCategory:  &data.SlugCategory,
		ImageCategory: &data.ImageCategory,
		CreatedAt:     &data.CreatedAt,
		UpdatedAt:     &data.UpdatedAt,
	}
}

func (c *categoryGraphqlMapper) mapResponsesCategory(categories []*pb.CategoryResponse) []*model.CategoryResponse {
	var responses []*model.CategoryResponse
	for _, category := range categories {
		responses = append(responses, c.mapResponseCategory(category))
	}
	return responses
}

func (c *categoryGraphqlMapper) mapResponseCategoryDeleteAt(data *pb.CategoryResponseDeleteAt) *model.CategoryResponseDeleteAt {
	if data == nil {
		return nil
	}

	var deletedAt string

	if data.DeletedAt != nil {
		deletedAt = data.DeletedAt.Value
	}

	return &model.CategoryResponseDeleteAt{
		ID:            int32(data.Id),
		Name:          data.Name,
		Description:   &data.Description,
		SlugCategory:  &data.SlugCategory,
		ImageCategory: &data.ImageCategory,
		CreatedAt:     &data.CreatedAt,
		UpdatedAt:     &data.UpdatedAt,
		DeletedAt:     &deletedAt,
	}
}

func (c *categoryGraphqlMapper) mapResponsesCategoryDeleteAt(categories []*pb.CategoryResponseDeleteAt) []*model.CategoryResponseDeleteAt {
	var responses []*model.CategoryResponseDeleteAt
	for _, category := range categories {
		responses = append(responses, c.mapResponseCategoryDeleteAt(category))
	}
	return responses
}

func (c *categoryGraphqlMapper) mapResponseCategoryMonthlyPrice(category *pb.CategoryMonthPriceResponse) *model.CategoryMonthPriceResponse {
	if category == nil {
		return nil
	}
	return &model.CategoryMonthPriceResponse{
		Month:        category.Month,
		CategoryID:   int32(category.CategoryId),
		CategoryName: category.CategoryName,
		OrderCount:   int32(category.OrderCount),
		ItemsSold:    int32(category.ItemsSold),
		TotalRevenue: int32(category.TotalRevenue),
	}
}

func (c *categoryGraphqlMapper) mapResponsesCategoryMonthlyPrice(categories []*pb.CategoryMonthPriceResponse) []*model.CategoryMonthPriceResponse {
	var responses []*model.CategoryMonthPriceResponse
	for _, category := range categories {
		responses = append(responses, c.mapResponseCategoryMonthlyPrice(category))
	}
	return responses
}

func (c *categoryGraphqlMapper) mapResponseCategoryYearlyPrice(category *pb.CategoryYearPriceResponse) *model.CategoryYearPriceResponse {
	if category == nil {
		return nil
	}
	return &model.CategoryYearPriceResponse{
		Year:         category.Year,
		CategoryID:   int32(category.CategoryId),
		CategoryName: category.CategoryName,
		OrderCount:   int32(category.OrderCount),
		ItemsSold:    int32(category.ItemsSold),
		TotalRevenue: int32(category.TotalRevenue),
	}
}

func (c *categoryGraphqlMapper) mapResponsesCategoryYearlyPrice(categories []*pb.CategoryYearPriceResponse) []*model.CategoryYearPriceResponse {
	var responses []*model.CategoryYearPriceResponse
	for _, category := range categories {
		responses = append(responses, c.mapResponseCategoryYearlyPrice(category))
	}
	return responses
}

func (c *categoryGraphqlMapper) mapResponseCategoryMonthlyTotalPrice(category *pb.CategoriesMonthlyTotalPriceResponse) *model.CategoriesMonthlyTotalPriceResponse {
	if category == nil {
		return nil
	}
	return &model.CategoriesMonthlyTotalPriceResponse{
		Month:        category.Month,
		Year:         category.Year,
		TotalRevenue: int32(category.TotalRevenue),
	}
}

func (c *categoryGraphqlMapper) mapResponsesCategoryMonthlyTotalPrice(categories []*pb.CategoriesMonthlyTotalPriceResponse) []*model.CategoriesMonthlyTotalPriceResponse {
	var responses []*model.CategoriesMonthlyTotalPriceResponse
	for _, category := range categories {
		responses = append(responses, c.mapResponseCategoryMonthlyTotalPrice(category))
	}
	return responses
}

func (c *categoryGraphqlMapper) mapResponseCategoryYearlyTotalPrice(category *pb.CategoriesYearlyTotalPriceResponse) *model.CategoriesYearlyTotalPriceResponse {
	if category == nil {
		return nil
	}
	return &model.CategoriesYearlyTotalPriceResponse{
		Year:         category.Year,
		TotalRevenue: int32(category.TotalRevenue),
	}
}

func (c *categoryGraphqlMapper) mapResponsesCategoryYearlyTotalPrice(categories []*pb.CategoriesYearlyTotalPriceResponse) []*model.CategoriesYearlyTotalPriceResponse {
	var responses []*model.CategoriesYearlyTotalPriceResponse
	for _, category := range categories {
		responses = append(responses, c.mapResponseCategoryYearlyTotalPrice(category))
	}
	return responses
}
