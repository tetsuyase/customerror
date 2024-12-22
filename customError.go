package customError

type CustomError struct {
	code int
	error
}

func New(code int, e error) error {
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
