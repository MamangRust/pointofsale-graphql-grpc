package graphql

import (
	"github.com/MamangRust/pointofsale-graphql-grpc/internal/model"
	"github.com/MamangRust/pointofsale-graphql-grpc/internal/pb"
)

type transactionGraphqlMapper struct {
}

func NewTransactionGraphqlMapper() *transactionGraphqlMapper {
	return &transactionGraphqlMapper{}
}

func (t *transactionGraphqlMapper) ToGraphqlResponseTransaction(res *pb.ApiResponseTransaction) *model.APIResponseTransaction {
	return &model.APIResponseTransaction{
		Status:  res.Status,
		Message: res.Message,
		Data:    t.mapResponseTransaction(res.Data),
	}
}

func (t *transactionGraphqlMapper) ToGraphqlResponsesTransaction(res *pb.ApiResponsesTransaction) *model.APIResponsesTransaction {
	return &model.APIResponsesTransaction{
		Status:  res.Status,
		Message: res.Message,
		Data:    t.mapResponsesTransaction(res.Data),
	}
}

func (t *transactionGraphqlMapper) ToGraphqlResponseTransactionDeleteAt(res *pb.ApiResponseTransactionDeleteAt) *model.APIResponseTransactionDeleteAt {
	return &model.APIResponseTransactionDeleteAt{
		Status:  res.Status,
		Message: res.Message,
		Data:    t.mapResponseTransactionDeleteAt(res.Data),
	}
}

func (t *transactionGraphqlMapper) ToGraphqlResponseTransactionDelete(res *pb.ApiResponseTransactionDelete) *model.APIResponseTransactionDelete {
	return &model.APIResponseTransactionDelete{
		Status:  res.Status,
		Message: res.Message,
	}
}

func (t *transactionGraphqlMapper) ToGraphqlResponseTransactionAll(res *pb.ApiResponseTransactionAll) *model.APIResponseTransactionAll {
	return &model.APIResponseTransactionAll{
		Status:  res.Status,
		Message: res.Message,
	}
}

func (t *transactionGraphqlMapper) ToGraphqlResponsePaginationTransaction(res *pb.ApiResponsePaginationTransaction) *model.APIResponsePaginationTransaction {
	return &model.APIResponsePaginationTransaction{
		Status:     res.Status,
		Message:    res.Message,
		Data:       t.mapResponsesTransaction(res.Data),
		Pagination: mapPaginationMeta(res.Pagination),
	}
}

func (t *transactionGraphqlMapper) ToGraphqlResponsePaginationTransactionDeleteAt(res *pb.ApiResponsePaginationTransactionDeleteAt) *model.APIResponsePaginationTransactionDeleteAt {
	return &model.APIResponsePaginationTransactionDeleteAt{
		Status:     res.Status,
		Message:    res.Message,
		Data:       t.mapResponsesTransactionDeleteAt(res.Data),
		Pagination: mapPaginationMeta(res.Pagination),
	}
}

func (t *transactionGraphqlMapper) ToGraphqlResponseMonthAmountSuccess(res *pb.ApiResponseTransactionMonthAmountSuccess) *model.APIResponseTransactionMonthAmountSuccess {
	return &model.APIResponseTransactionMonthAmountSuccess{
		Status:  res.Status,
		Message: res.Message,
		Data:    t.mapResponsesTransactionMonthlyAmountSuccess(res.Data),
	}
}

func (t *transactionGraphqlMapper) ToGraphqlResponseYearAmountSuccess(res *pb.ApiResponseTransactionYearAmountSuccess) *model.APIResponseTransactionYearAmountSuccess {
	return &model.APIResponseTransactionYearAmountSuccess{
		Status:  res.Status,
		Message: res.Message,
		Data:    t.mapResponsesTransactionYearlyAmountSuccess(res.Data),
	}
}

func (t *transactionGraphqlMapper) ToGraphqlResponseMonthAmountFailed(res *pb.ApiResponseTransactionMonthAmountFailed) *model.APIResponseTransactionMonthAmountFailed {
	return &model.APIResponseTransactionMonthAmountFailed{
		Status:  res.Status,
		Message: res.Message,
		Data:    t.mapResponsesTransactionMonthlyAmountFailed(res.Data),
	}
}

func (t *transactionGraphqlMapper) ToGraphqlResponseYearAmountFailed(res *pb.ApiResponseTransactionYearAmountFailed) *model.APIResponseTransactionYearAmountFailed {
	return &model.APIResponseTransactionYearAmountFailed{
		Status:  res.Status,
		Message: res.Message,
		Data:    t.mapResponsesTransactionYearlyAmountFailed(res.Data),
	}
}

func (t *transactionGraphqlMapper) ToGraphqlResponseMonthMethod(res *pb.ApiResponseTransactionMonthPaymentMethod) *model.APIResponseTransactionMonthPaymentMethod {
	return &model.APIResponseTransactionMonthPaymentMethod{
		Status:  res.Status,
		Message: res.Message,
		Data:    t.mapResponsesTransactionMonthlyMethod(res.Data),
	}
}

func (t *transactionGraphqlMapper) ToGraphqlResponseYearMethod(res *pb.ApiResponseTransactionYearPaymentmethod) *model.APIResponseTransactionYearPaymentMethod {
	return &model.APIResponseTransactionYearPaymentMethod{
		Status:  res.Status,
		Message: res.Message,
		Data:    t.mapResponsesTransactionYearlyMethod(res.Data),
	}
}

func (t *transactionGraphqlMapper) mapResponseTransaction(transaction *pb.TransactionResponse) *model.TransactionResponse {
	id := int32(transaction.Id)
	orderID := int32(transaction.OrderId)
	merchantID := int32(transaction.MerchantId)
	amount := int32(transaction.Amount)
	changeAmount := int32(transaction.ChangeAmount)

	return &model.TransactionResponse{
		ID:            id,
		OrderID:       orderID,
		MerchantID:    merchantID,
		PaymentMethod: transaction.PaymentMethod,
		Amount:        amount,
		ChangeAmount:  &changeAmount,
		PaymentStatus: transaction.PaymentStatus,
		CreatedAt:     transaction.CreatedAt,
		UpdatedAt:     transaction.UpdatedAt,
	}
}

func (t *transactionGraphqlMapper) mapResponsesTransaction(transactions []*pb.TransactionResponse) []*model.TransactionResponse {
	var mapped []*model.TransactionResponse
	for _, tx := range transactions {
		mapped = append(mapped, t.mapResponseTransaction(tx))
	}
	return mapped
}

func (t *transactionGraphqlMapper) mapResponseTransactionDeleteAt(transaction *pb.TransactionResponseDeleteAt) *model.TransactionResponseDeleteAt {
	id := int32(transaction.Id)
	orderID := int32(transaction.OrderId)
	merchantID := int32(transaction.MerchantId)
	amount := int32(transaction.Amount)
	changeAmount := int32(transaction.ChangeAmount)

	var deletedAt *string
	if transaction.DeletedAt != nil {
		deletedAt = &transaction.DeletedAt.Value
	}

	return &model.TransactionResponseDeleteAt{
		ID:            id,
		OrderID:       orderID,
		MerchantID:    merchantID,
		PaymentMethod: transaction.PaymentMethod,
		Amount:        amount,
		ChangeAmount:  &changeAmount,
		PaymentStatus: transaction.PaymentStatus,
		CreatedAt:     transaction.CreatedAt,
		UpdatedAt:     transaction.UpdatedAt,
		DeletedAt:     deletedAt,
	}
}

func (t *transactionGraphqlMapper) mapResponsesTransactionDeleteAt(transactions []*pb.TransactionResponseDeleteAt) []*model.TransactionResponseDeleteAt {
	var mapped []*model.TransactionResponseDeleteAt
	for _, tx := range transactions {
		mapped = append(mapped, t.mapResponseTransactionDeleteAt(tx))
	}
	return mapped
}

func (t *transactionGraphqlMapper) mapResponseTransactionMonthAmountSuccess(row *pb.TransactionMonthlyAmountSuccess) *model.TransactionMonthlyAmountSuccess {
	totalSuccess := int32(row.TotalSuccess)
	totalAmount := int32(row.TotalAmount)

	return &model.TransactionMonthlyAmountSuccess{
		Year:         row.Year,
		Month:        row.Month,
		TotalSuccess: totalSuccess,
		TotalAmount:  totalAmount,
	}
}

func (t *transactionGraphqlMapper) mapResponsesTransactionMonthlyAmountSuccess(rows []*pb.TransactionMonthlyAmountSuccess) []*model.TransactionMonthlyAmountSuccess {
	var mapped []*model.TransactionMonthlyAmountSuccess
	for _, row := range rows {
		mapped = append(mapped, t.mapResponseTransactionMonthAmountSuccess(row))
	}
	return mapped
}

func (t *transactionGraphqlMapper) mapResponseTransactionYearAmountSuccess(row *pb.TransactionYearlyAmountSuccess) *model.TransactionYearlyAmountSuccess {
	totalSuccess := int32(row.TotalSuccess)
	totalAmount := int32(row.TotalAmount)

	return &model.TransactionYearlyAmountSuccess{
		Year:         row.Year,
		TotalSuccess: totalSuccess,
		TotalAmount:  totalAmount,
	}
}

func (t *transactionGraphqlMapper) mapResponsesTransactionYearlyAmountSuccess(rows []*pb.TransactionYearlyAmountSuccess) []*model.TransactionYearlyAmountSuccess {
	var mapped []*model.TransactionYearlyAmountSuccess
	for _, row := range rows {
		mapped = append(mapped, t.mapResponseTransactionYearAmountSuccess(row))
	}
	return mapped
}

func (t *transactionGraphqlMapper) mapResponseTransactionMonthAmountFailed(row *pb.TransactionMonthlyAmountFailed) *model.TransactionMonthlyAmountFailed {
	totalFailed := int32(row.TotalFailed)
	totalAmount := int32(row.TotalAmount)

	return &model.TransactionMonthlyAmountFailed{
		Year:        row.Year,
		Month:       row.Month,
		TotalFailed: totalFailed,
		TotalAmount: totalAmount,
	}
}

func (t *transactionGraphqlMapper) mapResponsesTransactionMonthlyAmountFailed(rows []*pb.TransactionMonthlyAmountFailed) []*model.TransactionMonthlyAmountFailed {
	var mapped []*model.TransactionMonthlyAmountFailed
	for _, row := range rows {
		mapped = append(mapped, t.mapResponseTransactionMonthAmountFailed(row))
	}
	return mapped
}

func (t *transactionGraphqlMapper) mapResponseTransactionYearAmountFailed(row *pb.TransactionYearlyAmountFailed) *model.TransactionYearlyAmountFailed {
	totalFailed := int32(row.TotalFailed)
	totalAmount := int32(row.TotalAmount)

	return &model.TransactionYearlyAmountFailed{
		Year:        row.Year,
		TotalFailed: totalFailed,
		TotalAmount: totalAmount,
	}
}

func (t *transactionGraphqlMapper) mapResponsesTransactionYearlyAmountFailed(rows []*pb.TransactionYearlyAmountFailed) []*model.TransactionYearlyAmountFailed {
	var mapped []*model.TransactionYearlyAmountFailed
	for _, row := range rows {
		mapped = append(mapped, t.mapResponseTransactionYearAmountFailed(row))
	}
	return mapped
}

func (t *transactionGraphqlMapper) mapResponseTransactionMonthMethod(row *pb.TransactionMonthlyMethod) *model.TransactionMonthlyMethod {
	totalTransactions := int32(row.TotalTransactions)
	totalAmount := int32(row.TotalAmount)

	return &model.TransactionMonthlyMethod{
		Month:             row.Month,
		PaymentMethod:     row.PaymentMethod,
		TotalTransactions: totalTransactions,
		TotalAmount:       totalAmount,
	}
}

func (t *transactionGraphqlMapper) mapResponsesTransactionMonthlyMethod(rows []*pb.TransactionMonthlyMethod) []*model.TransactionMonthlyMethod {
	var mapped []*model.TransactionMonthlyMethod
	for _, row := range rows {
		mapped = append(mapped, t.mapResponseTransactionMonthMethod(row))
	}
	return mapped
}

func (t *transactionGraphqlMapper) mapResponseTransactionYearMethod(row *pb.TransactionYearlyMethod) *model.TransactionYearlyMethod {
	totalTransactions := int32(row.TotalTransactions)
	totalAmount := int32(row.TotalAmount)

	return &model.TransactionYearlyMethod{
		Year:              row.Year,
		PaymentMethod:     row.PaymentMethod,
		TotalTransactions: totalTransactions,
		TotalAmount:       totalAmount,
	}
}

func (t *transactionGraphqlMapper) mapResponsesTransactionYearlyMethod(rows []*pb.TransactionYearlyMethod) []*model.TransactionYearlyMethod {
	var mapped []*model.TransactionYearlyMethod
	for _, row := range rows {
		mapped = append(mapped, t.mapResponseTransactionYearMethod(row))
	}
	return mapped
}
