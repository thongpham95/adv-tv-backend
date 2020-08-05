package responsehandler

import (
	"encoding/json"
	"net/http"
)

// HTTPResponse is custom response struct
type HTTPResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// SuccessResponse returns body as JSON
func (r *HTTPResponse) SuccessResponse(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(r)
}

// ErrorResponse returns body as JSON
func (r *HTTPResponse) ErrorResponse(w http.ResponseWriter, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(r)
}

// NewHTTPResponse return a new HTTPResponse
func NewHTTPResponse(isSuccess bool, message string, data interface{}) *HTTPResponse {
	return &HTTPResponse{
		Success: isSuccess,
		Message: message,
		Data:    data,
	}
}
