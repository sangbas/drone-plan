package entity

import "github.com/SawitProRecruitment/UserService/generated"

var (
	BadRequestError     = NewErrorResponse("Invalid input")
	InternalServerError = NewErrorResponse("Internal server error")
)

func NewErrorResponse(message string) *generated.ErrorResponse {
	return &generated.ErrorResponse{
		Message: message,
	}
}
