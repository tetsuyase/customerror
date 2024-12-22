package customError

type CustomError struct {
	code      int
	requestID int
	error
}

func New(code int, requestID int, e error) error {
	return &CustomError{
		code:      code,
		requestID: requestID,
		error:     e,
	}
}

func (a *CustomError) Error() string {
	return ""
}

func (a *CustomError) GetError() error {
	return a.error
}

func (a *CustomError) GetCode() int {
	return a.code
}
