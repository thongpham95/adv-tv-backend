package responsehandler

import (
	"encoding/json"
	"net/http"
)

// HTTPResponse is custom response struct
type HTTPResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
}

// SuccessResponse returns body as JSON
func (r *HTTPResponse) SuccessResponse(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(r)
}

// NewHTTPResponse return a new HTTPResponse
func NewHTTPResponse(data interface{}) *HTTPResponse {
	return &HTTPResponse{
		Success: true,
		Data:    data,
	}
}
