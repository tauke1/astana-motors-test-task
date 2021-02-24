package model

type errorResponse struct {
	Message string `json:"message"`
}

func NewErrorResponse(message string) *errorResponse {
	return &errorResponse{
		Message: message,
	}
}
