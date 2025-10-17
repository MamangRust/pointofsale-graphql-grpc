package errors

import (
	"encoding/json"

	"github.com/MamangRust/pointofsale-graphql-grpc/internal/pb"
)

func GrpcErrorToJson(err *pb.ErrorResponse) string {
	jsonData, _ := json.Marshal(err)
	return string(jsonData)
}
