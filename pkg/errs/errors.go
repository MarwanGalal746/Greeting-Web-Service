package errs

type ErrorResponse struct {
	Message   string `json:"message"`
	ErrorCode int    `json:"error_code"`
}

var (
	// ErrInternalServer application error code not HTTP error code
	// 8000, 8001 etc. mens server errors
	// 7000, 7001 etc. means the request contains bad syntax or cannot be fulfilled
	ErrInternalServer   = ErrorResponse{Message: `Oops, something went wrong!`, ErrorCode: 8000}
	ErrInvalidJson      = ErrorResponse{Message: `Invalid JSON`, ErrorCode: 7000}
	ErrMethodNotAllowed = ErrorResponse{Message: `This method is not allowed, POST method only is allowed`, ErrorCode: 7001}
	ErrRouteNotFound    = ErrorResponse{Message: `This path is not found, "greet/" exists only`, ErrorCode: 7002}
)
