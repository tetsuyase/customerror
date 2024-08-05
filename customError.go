package customError

type CustomError struct {
	code int
	error
}

func New(e error, code int) error {
	return &CustomError{
		code:  code,
		error: e,
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
