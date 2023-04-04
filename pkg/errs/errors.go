package errs

type ErrorResponse struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

var (
	// ErrInternalServer application error code not HTTP status code
	// 8000, 8001 etc mens server errors
	// 7000, 7001 etc means the request contains bad syntax or cannot be fulfilled
	ErrInternalServer   = ErrorResponse{Message: `Oops, something went wrong!`, Status: 8000}
	ErrInvalidJson      = ErrorResponse{Message: `Invalid JSON`, Status: 7000}
	ErrMethodNotAllowed = ErrorResponse{Message: `This method is not allowed, POST method only is allowed`, Status: 7001}
	ErrRouteNotFound    = ErrorResponse{Message: `This path is not found, "greet/" exists only`, Status: 7002}
)
