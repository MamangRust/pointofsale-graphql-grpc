package graphql

import (
	"github.com/MamangRust/pointofsale-graphql-grpc/internal/model"
	"github.com/MamangRust/pointofsale-graphql-grpc/internal/pb"
)

type AuthGraphqlMapper interface {
	ToGraphqlResponseLogin(resp *pb.ApiResponseLogin) *model.APIResponseLogin
	ToGraphqlResponseRegister(resp *pb.ApiResponseRegister) *model.APIResponseRegister
	ToGraphqlResponseRefreshToken(resp *pb.ApiResponseRefreshToken) *model.APIResponseRefreshToken
	ToGraphqlResponseGetMe(resp *pb.ApiResponseGetMe) *model.APIResponseGetMe
}

type UserGraphqlMapper interface {
	ToGraphqlResponseUser(resp *pb.ApiResponseUser) *model.APIResponseUserResponse
	ToGraphqlResponseUserDeleteAt(resp *pb.ApiResponseUserDeleteAt) *model.APIResponseUserResponseDeleteAt
	ToGraphqlResponseUsers(resp *pb.ApiResponsesUser) *model.APIResponsesUser
	ToGraphqlResponseUserDelete(resp *pb.ApiResponseUserDelete) *model.APIResponseUserDelete
	ToGraphqlResponseUserAll(resp *pb.ApiResponseUserAll) *model.APIResponseUserAll
	ToGraphqlResponsePaginationUser(resp *pb.ApiResponsePaginationUser) *model.APIResponsePaginationUser
	ToGraphqlResponsePaginationUserDeleteAt(resp *pb.ApiResponsePaginationUserDeleteAt) *model.APIResponsePaginationUserDeleteAt
}

type RoleGraphqlMapper interface {
	ToGraphqlResponseRole(resp *pb.ApiResponseRole) *model.APIResponseRole
	ToGraphqlResponseRoleDeleteAt(resp *pb.ApiResponseRoleDeleteAt) *model.APIResponseRoleDeleteAt
	ToGraphqlResponsesRole(resp *pb.ApiResponsesRole) *model.APIResponsesRole
	ToGraphqlResponseDelete(resp *pb.ApiResponseRoleDelete) *model.APIResponseRoleDelete
	ToGraphqlResponseAll(resp *pb.ApiResponseRoleAll) *model.APIResponseRoleAll
	ToGraphqlResponsePaginationRole(resp *pb.ApiResponsePaginationRole) *model.APIResponsePaginationRole
	ToGraphqlResponsePaginationRoleDeleteAt(resp *pb.ApiResponsePaginationRoleDeleteAt) *model.APIResponsePaginationRoleDeleteAt
}

type CashierGraphqlMapper interface {
	ToGraphqlResponseCashier(resp *pb.ApiResponseCashier) *model.APIResponseCashier
	ToGraphqlResponsesCashier(resp *pb.ApiResponsesCashier) *model.APIResponsesCashier
	ToGraphqlResponseCashierDeleteAt(resp *pb.ApiResponseCashierDeleteAt) *model.APIResponseCashierDeleteAt
	ToGraphqlResponseCashierDelete(resp *pb.ApiResponseCashierDelete) *model.APIResponseCashierDelete
	ToGraphqlResponseCashierAll(resp *pb.ApiResponseCashierAll) *model.APIResponseCashierAll
	ToGraphqlResponsePaginationCashier(resp *pb.ApiResponsePaginationCashier) *model.APIResponsePaginationCashier
	ToGraphqlResponsePaginationCashierDeleteAt(resp *pb.ApiResponsePaginationCashierDeleteAt) *model.APIResponsePaginationCashierDeleteAt
	ToGraphqlResponseMonthlySales(resp *pb.ApiResponseCashierMonthSales) *model.APIResponseCashierMonthSales
	ToGraphqlResponseYearlySales(resp *pb.ApiResponseCashierYearSales) *model.APIResponseCashierYearSales
	ToGraphqlResponseMonthlyTotalSales(resp *pb.ApiResponseCashierMonthlyTotalSales) *model.APIResponseCashierMonthlyTotalSales
	ToGraphqlResponseYearlyTotalSales(resp *pb.ApiResponseCashierYearlyTotalSales) *model.APIResponseCashierYearlyTotalSales
}

type CategoryGraphqlMapper interface {
	ToGraphqlResponseCategory(resp *pb.ApiResponseCategory) *model.APIResponseCategory
	ToGraphqlResponseCategoryDeleteAt(resp *pb.ApiResponseCategoryDeleteAt) *model.APIResponseCategoryDeleteAt
	ToGraphqlResponsesCategory(resp *pb.ApiResponsesCategory) *model.APIResponsesCategory
	ToGraphqlResponseCategoryDelete(resp *pb.ApiResponseCategoryDelete) *model.APIResponseCategoryDelete
	ToGraphqlResponseCategoryAll(resp *pb.ApiResponseCategoryAll) *model.APIResponseCategoryAll
	ToGraphqlResponsePaginationCategory(resp *pb.ApiResponsePaginationCategory) *model.APIResponsePaginationCategory
	ToGraphqlResponsePaginationCategoryDeleteAt(resp *pb.ApiResponsePaginationCategoryDeleteAt) *model.APIResponsePaginationCategoryDeleteAt
	ToGraphqlResponseCategoryMonthlyPrice(resp *pb.ApiResponseCategoryMonthPrice) *model.APIResponseCategoryMonthPrice
	ToGraphqlResponseCategoryYearlyPrice(resp *pb.ApiResponseCategoryYearPrice) *model.APIResponseCategoryYearPrice
	ToGraphqlResponseCategoryMonthlyTotalPrice(resp *pb.ApiResponseCategoryMonthlyTotalPrice) *model.APIResponseCategoryMonthlyTotalPrice
	ToGraphqlResponseCategoryYearlyTotalPrice(resp *pb.ApiResponseCategoryYearlyTotalPrice) *model.APIResponseCategoryYearlyTotalPrice
}

type MerchantGraphqlMapper interface {
	ToGraphqlResponseMerchant(resp *pb.ApiResponseMerchant) *model.APIResponseMerchant
	ToGraphqlResponsesMerchant(resp *pb.ApiResponsesMerchant) *model.APIResponsesMerchant
	ToGraphqlResponseMerchantDeleteAt(resp *pb.ApiResponseMerchantDeleteAt) *model.APIResponseMerchantDeleteAt
	ToGraphqlResponseMerchantAll(resp *pb.ApiResponseMerchantAll) *model.APIResponseMerchantAll
	ToGraphqlResponseMerchantDelete(resp *pb.ApiResponseMerchantDelete) *model.APIResponseMerchantDelete
	ToGraphqlResponsePaginationMerchantDeleteAt(resp *pb.ApiResponsePaginationMerchantDeleteAt) *model.APIResponsePaginationMerchantDeleteAt
	ToGraphqlResponsePaginationMerchant(resp *pb.ApiResponsePaginationMerchant) *model.APIResponsePaginationMerchant
}

type OrderItemGraphqlMapper interface {
	ToGraphqlResponseOrderItem(resp *pb.ApiResponseOrderItem) *model.APIResponseOrderItem
	ToGraphqlResponsesOrderItem(resp *pb.ApiResponsesOrderItem) *model.APIResponsesOrderItem
	ToGrapqhlResponseOrderItemDelete(resp *pb.ApiResponseOrderItemDelete) *model.APIResponseOrderItemDelete
	ToGrapqhlResponseOrderItemAll(resp *pb.ApiResponseOrderItemAll) *model.APIResponseOrderItemAll
	ToGraphqlResponsePaginationOrderItem(resp *pb.ApiResponsePaginationOrderItem) *model.APIResponsePaginationOrderItem
	ToGraphqlResponsePaginationOrderItemDeleteAt(resp *pb.ApiResponsePaginationOrderItemDeleteAt) *model.APIResponsePaginationOrderItemDeleteAt
}

type OrderGraphqlMapper interface {
	ToGraphqlResponseOrder(resp *pb.ApiResponseOrder) *model.APIResponseOrder
	ToGraphqlResponsesOrder(resp *pb.ApiResponsesOrder) *model.APIResponsesOrder
	ToGraphqlResponseOrderDeleteAt(resp *pb.ApiResponseOrderDeleteAt) *model.APIResponseOrderDeleteAt
	ToGraphqlResponseOrderDelete(resp *pb.ApiResponseOrderDelete) *model.APIResponseOrderDelete
	ToGraphqlResponseOrderAll(resp *pb.ApiResponseOrderAll) *model.APIResponseOrderAll
	ToGraphqlResponsePaginationOrder(resp *pb.ApiResponsePaginationOrder) *model.APIResponsePaginationOrder
	ToGraphqlResponsePaginationOrderDeleteAt(resp *pb.ApiResponsePaginationOrderDeleteAt) *model.APIResponsePaginationOrderDeleteAt
	ToGraphqlResponseMonthlyRevenue(resp *pb.ApiResponseOrderMonthly) *model.APIResponseOrderMonthly
	ToGraphqlResponseYearlyRevenue(resp *pb.ApiResponseOrderYearly) *model.APIResponseOrderYearly
	ToGraphqlResponseMonthlyTotalRevenue(resp *pb.ApiResponseOrderMonthlyTotalRevenue) *model.APIResponseOrderMonthlyTotalRevenue
	ToGraphqlResponseYearlyTotalRevenue(resp *pb.ApiResponseOrderYearlyTotalRevenue) *model.APIResponseOrderYearlyTotalRevenue
}

type ProductGraphqlMapper interface {
	ToGraphqlResponseProduct(resp *pb.ApiResponseProduct) *model.APIResponseProduct
	ToGraphqlResponsesProduct(resp *pb.ApiResponsesProduct) *model.APIResponsesProduct
	ToGraphqlResponseProductDeleteAt(resp *pb.ApiResponseProductDeleteAt) *model.APIResponseProductDeleteAt
	ToGraphqlResponseProductDelete(resp *pb.ApiResponseProductDelete) *model.APIResponseProductDelete
	ToGraphqlResponseProductAll(resp *pb.ApiResponseProductAll) *model.APIResponseProductAll
	ToGraphqlResponsePaginationProduct(resp *pb.ApiResponsePaginationProduct) *model.APIResponsePaginationProduct
	ToGraphqlResponsePaginationProductDeleteAt(resp *pb.ApiResponsePaginationProductDeleteAt) *model.APIResponsePaginationProductDeleteAt
}

type TransactionGraphqlMapper interface {
	ToGraphqlResponseTransaction(resp *pb.ApiResponseTransaction) *model.APIResponseTransaction
	ToGraphqlResponsesTransaction(resp *pb.ApiResponsesTransaction) *model.APIResponsesTransaction
	ToGraphqlResponseTransactionDeleteAt(resp *pb.ApiResponseTransactionDeleteAt) *model.APIResponseTransactionDeleteAt
	ToGraphqlResponseTransactionDelete(resp *pb.ApiResponseTransactionDelete) *model.APIResponseTransactionDelete
	ToGraphqlResponseTransactionAll(resp *pb.ApiResponseTransactionAll) *model.APIResponseTransactionAll
	ToGraphqlResponsePaginationTransaction(resp *pb.ApiResponsePaginationTransaction) *model.APIResponsePaginationTransaction
	ToGraphqlResponsePaginationTransactionDeleteAt(resp *pb.ApiResponsePaginationTransactionDeleteAt) *model.APIResponsePaginationTransactionDeleteAt
	ToGraphqlResponseMonthAmountSuccess(resp *pb.ApiResponseTransactionMonthAmountSuccess) *model.APIResponseTransactionMonthAmountSuccess
	ToGraphqlResponseYearAmountSuccess(resp *pb.ApiResponseTransactionYearAmountSuccess) *model.APIResponseTransactionYearAmountSuccess
	ToGraphqlResponseMonthAmountFailed(resp *pb.ApiResponseTransactionMonthAmountFailed) *model.APIResponseTransactionMonthAmountFailed
	ToGraphqlResponseYearAmountFailed(resp *pb.ApiResponseTransactionYearAmountFailed) *model.APIResponseTransactionYearAmountFailed
	ToGraphqlResponseMonthMethod(resp *pb.ApiResponseTransactionMonthPaymentMethod) *model.APIResponseTransactionMonthPaymentMethod
	ToGraphqlResponseYearMethod(resp *pb.ApiResponseTransactionYearPaymentmethod) *model.APIResponseTransactionYearPaymentMethod
}
