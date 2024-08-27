package domain



type APIResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   interface{} `json:"error,omitempty"`
}

// NewSuccessResponse creates a success response with the given message and data.
func NewSuccessResponse(message string, data interface{}) *APIResponse {
	return &APIResponse{
		Status:  "success",
		Message: message,
		Data:    data,
	}
}

// NewErrorResponse creates an error response with the given message and error code.
func NewErrorResponse(message string, err error) *APIResponse {
	return &APIResponse{
		Status:  "error",
		Message: message,
		Error: map[string]interface{}{
			"code":    GetStatusCode(err),
			"details": err.Error(),
		},
	}
}
