package helper

type Error struct {
	code    int
	message string
	cause   error
}

func newError(code int, msg string) *Error {
	return &Error{code: code, message: msg}
}

func (e *Error) Error() string {
	return e.message
}

func (e *Error) Code() int {
	return e.code
}

func (e *Error) Message() string {
	return e.message
}

func (e *Error) Unwrap() error {
	return e.cause
}

func (e *Error) WithCause(err error) error {
	e.cause = err
	return e
}
