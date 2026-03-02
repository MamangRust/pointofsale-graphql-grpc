package graphql

type GraphqlMapper struct {
	AuthGraphqlMapper
	RoleGraphqlMapper
	UserGraphqlMapper
	CashierGraphqlMapper
	CategoryGraphqlMapper
	MerchantGraphqlMapper
	OrderItemGraphqlMapper
	OrderGraphqlMapper
	ProductGraphqlMapper
	TransactionGraphqlMapper
}

func NewGraphqlMapper() *GraphqlMapper {
	return &GraphqlMapper{
		AuthGraphqlMapper:        NewAuthResponseMapper(),
		RoleGraphqlMapper:        NewRoleResponseMapper(),
		UserGraphqlMapper:        NewUserResponseMapper(),
		CashierGraphqlMapper:     NewCashierGraphqlMapper(),
		CategoryGraphqlMapper:    NewCategoryGraphqlMapper(),
		MerchantGraphqlMapper:    NewMerchantGraphqlMapper(),
		OrderItemGraphqlMapper:   NewOrderItemGraphqlMapper(),
		OrderGraphqlMapper:       NewOrderGraphqlMapper(),
		ProductGraphqlMapper:     NewProductGraphqlMapper(),
		TransactionGraphqlMapper: NewTransactionGraphqlMapper(),
	}
}
