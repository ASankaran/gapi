package errors

type APIError struct {
    Status     int      `json:"status,omitempty"`
    Message    string   `json:"message,omitempty"`
}
