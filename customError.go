package customError

type customError struct {
	code int
	error
}

func New(code int, e error) error {
	return &customError{
		code:  code,
		error: e,
	}
}

func (a *customError) Error() string {
	return ""
}

func (a *customError) GetError() error {
	return a.error
}

func (a *customError) GetCode() int {
	return a.code
}
