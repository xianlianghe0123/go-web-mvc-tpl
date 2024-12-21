package helper

var (
	// common errors
	ErrUnknownError        = newError(10001, "Unknown Error")
	ErrInternalServerError = newError(10002, "Internal Server Error")
	ErrBadRequest          = newError(10003, "Bad Request")
	ErrUnauthorized        = newError(10004, "Unauthorized")
	// more biz errors
	UserNotFound = newError(20001, "User not found")
)
