package graphql

import (
	"github.com/MamangRust/pointofsale-graphql-grpc/internal/model"
	"github.com/MamangRust/pointofsale-graphql-grpc/internal/pb"
)

type cashierGraphqlMapper struct {
}

func NewCashierGraphqlMapper() *cashierGraphqlMapper {
	return &cashierGraphqlMapper{}
}

func (c *cashierGraphqlMapper) ToGraphqlResponseCashier(res *pb.ApiResponseCashier) *model.APIResponseCashier {
	return &model.APIResponseCashier{
		Status:  res.Status,
		Message: res.Message,
		Data:    c.mapResponseCashier(res.Data),
	}
}

func (c *cashierGraphqlMapper) ToGraphqlResponsesCashier(res *pb.ApiResponsesCashier) *model.APIResponsesCashier {
	return &model.APIResponsesCashier{
		Status:  res.Status,
		Message: res.Message,
		Data:    c.mapResponsesCashier(res.Data),
	}
}

func (c *cashierGraphqlMapper) ToGraphqlResponseCashierDeleteAt(res *pb.ApiResponseCashierDeleteAt) *model.APIResponseCashierDeleteAt {
	return &model.APIResponseCashierDeleteAt{
		Status:  res.Status,
		Message: res.Message,
		Data:    c.mapResponseCashierDeleteAt(res.Data),
	}
}

func (c *cashierGraphqlMapper) ToGraphqlResponseCashierDelete(res *pb.ApiResponseCashierDelete) *model.APIResponseCashierDelete {
	return &model.APIResponseCashierDelete{
		Status:  res.Status,
		Message: res.Message,
	}
}

func (c *cashierGraphqlMapper) ToGraphqlResponseCashierAll(res *pb.ApiResponseCashierAll) *model.APIResponseCashierAll {
	return &model.APIResponseCashierAll{
		Status:  res.Status,
		Message: res.Message,
	}
}

func (c *cashierGraphqlMapper) ToGraphqlResponsePaginationCashier(res *pb.ApiResponsePaginationCashier) *model.APIResponsePaginationCashier {
	return &model.APIResponsePaginationCashier{
		Status:     res.Status,
		Message:    res.Message,
		Data:       c.mapResponsesCashier(res.Data),
		Pagination: mapPaginationMeta(res.Pagination),
	}
}

func (c *cashierGraphqlMapper) ToGraphqlResponsePaginationCashierDeleteAt(res *pb.ApiResponsePaginationCashierDeleteAt) *model.APIResponsePaginationCashierDeleteAt {
	return &model.APIResponsePaginationCashierDeleteAt{
		Status:     res.Status,
		Message:    res.Message,
		Data:       c.mapResponsesCashierDeleteAt(res.Data),
		Pagination: mapPaginationMeta(res.Pagination),
	}
}

func (c *cashierGraphqlMapper) ToGraphqlResponseMonthlySales(res *pb.ApiResponseCashierMonthSales) *model.APIResponseCashierMonthSales {
	return &model.APIResponseCashierMonthSales{
		Status:  res.Status,
		Message: res.Message,
		Data:    c.mapResponsesCashierMonthlySale(res.Data),
	}
}

func (c *cashierGraphqlMapper) ToGraphqlResponseYearlySales(res *pb.ApiResponseCashierYearSales) *model.APIResponseCashierYearSales {
	return &model.APIResponseCashierYearSales{
		Status:  res.Status,
		Message: res.Message,
		Data:    c.mapResponsesCashierYearlySale(res.Data),
	}
}

func (c *cashierGraphqlMapper) ToGraphqlResponseMonthlyTotalSales(res *pb.ApiResponseCashierMonthlyTotalSales) *model.APIResponseCashierMonthlyTotalSales {
	return &model.APIResponseCashierMonthlyTotalSales{
		Status:  res.Status,
		Message: res.Message,
		Data:    c.mapResponsesCashierMonthlyTotalSale(res.Data),
	}
}

func (c *cashierGraphqlMapper) ToGraphqlResponseYearlyTotalSales(res *pb.ApiResponseCashierYearlyTotalSales) *model.APIResponseCashierYearlyTotalSales {
	return &model.APIResponseCashierYearlyTotalSales{
		Status:  res.Status,
		Message: res.Message,
		Data:    c.mapResponsesCashierYearlyTotalSale(res.Data),
	}
}

func (c *cashierGraphqlMapper) mapResponseCashier(cashier *pb.CashierResponse) *model.CashierResponse {
	if cashier == nil {
		return nil
	}
	return &model.CashierResponse{
		ID:         cashier.Id,
		MerchantID: cashier.MerchantId,
		Name:       cashier.Name,
		CreatedAt:  &cashier.CreatedAt,
		UpdatedAt:  &cashier.UpdatedAt,
	}
}

func (c *cashierGraphqlMapper) mapResponsesCashier(cashiers []*pb.CashierResponse) []*model.CashierResponse {
	var responses []*model.CashierResponse
	for _, cashier := range cashiers {
		responses = append(responses, c.mapResponseCashier(cashier))
	}
	return responses
}

func (c *cashierGraphqlMapper) mapResponseCashierDeleteAt(cashier *pb.CashierResponseDeleteAt) *model.CashierResponseDeleteAt {
	if cashier == nil {
		return nil
	}

	var deletedAt string

	if cashier.DeletedAt != nil {
		deletedAt = cashier.DeletedAt.Value
	}

	return &model.CashierResponseDeleteAt{
		ID:         cashier.Id,
		MerchantID: cashier.MerchantId,
		Name:       cashier.Name,
		CreatedAt:  &cashier.CreatedAt,
		UpdatedAt:  &cashier.UpdatedAt,
		DeletedAt:  &deletedAt,
	}
}

func (c *cashierGraphqlMapper) mapResponsesCashierDeleteAt(cashiers []*pb.CashierResponseDeleteAt) []*model.CashierResponseDeleteAt {
	var responses []*model.CashierResponseDeleteAt
	for _, cashier := range cashiers {
		responses = append(responses, c.mapResponseCashierDeleteAt(cashier))
	}
	return responses
}

func (c *cashierGraphqlMapper) mapResponseCashierMonthlySale(cashier *pb.CashierResponseMonthSales) *model.CashierResponseMonthSales {
	if cashier == nil {
		return nil
	}
	return &model.CashierResponseMonthSales{
		Month:       cashier.Month,
		CashierID:   cashier.CashierId,
		CashierName: cashier.CashierName,
		OrderCount:  cashier.OrderCount,
		TotalSales:  cashier.TotalSales,
	}
}

func (c *cashierGraphqlMapper) mapResponsesCashierMonthlySale(cashiers []*pb.CashierResponseMonthSales) []*model.CashierResponseMonthSales {
	var responses []*model.CashierResponseMonthSales
	for _, cashier := range cashiers {
		responses = append(responses, c.mapResponseCashierMonthlySale(cashier))
	}
	return responses
}

func (c *cashierGraphqlMapper) mapResponseCashierYearlySale(cashier *pb.CashierResponseYearSales) *model.CashierResponseYearSales {
	if cashier == nil {
		return nil
	}
	return &model.CashierResponseYearSales{
		Year:        cashier.Year,
		CashierID:   cashier.CashierId,
		CashierName: cashier.CashierName,
		OrderCount:  cashier.OrderCount,
		TotalSales:  cashier.TotalSales,
	}
}

func (c *cashierGraphqlMapper) mapResponsesCashierYearlySale(cashiers []*pb.CashierResponseYearSales) []*model.CashierResponseYearSales {
	var responses []*model.CashierResponseYearSales
	for _, cashier := range cashiers {
		responses = append(responses, c.mapResponseCashierYearlySale(cashier))
	}
	return responses
}

func (c *cashierGraphqlMapper) mapResponseCashierMonthlyTotalSale(cashier *pb.CashierResponseMonthTotalSales) *model.CashierResponseMonthTotalSales {
	if cashier == nil {
		return nil
	}
	return &model.CashierResponseMonthTotalSales{
		Year:       cashier.Year,
		Month:      cashier.Month,
		TotalSales: cashier.TotalSales,
	}
}

func (c *cashierGraphqlMapper) mapResponsesCashierMonthlyTotalSale(cashiers []*pb.CashierResponseMonthTotalSales) []*model.CashierResponseMonthTotalSales {
	var responses []*model.CashierResponseMonthTotalSales
	for _, cashier := range cashiers {
		responses = append(responses, c.mapResponseCashierMonthlyTotalSale(cashier))
	}
	return responses
}

func (c *cashierGraphqlMapper) mapResponseCashierYearlyTotalSale(cashier *pb.CashierResponseYearTotalSales) *model.CashierResponseYearTotalSales {
	if cashier == nil {
		return nil
	}
	return &model.CashierResponseYearTotalSales{
		Year:       cashier.Year,
		TotalSales: cashier.TotalSales,
	}
}

func (c *cashierGraphqlMapper) mapResponsesCashierYearlyTotalSale(cashiers []*pb.CashierResponseYearTotalSales) []*model.CashierResponseYearTotalSales {
	var responses []*model.CashierResponseYearTotalSales
	for _, cashier := range cashiers {
		responses = append(responses, c.mapResponseCashierYearlyTotalSale(cashier))
	}
	return responses
}
