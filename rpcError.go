package rpcError

const (
	ParseError     = -32700
	MethodNotFound = -32601
	InvalidParams  = -32602
	InternalError  = -32602
)

type rpcError struct {
	code int
	error
}

func New(code int, e error) error {
	return &rpcError{
		code:  code,
		error: e,
	}
}

func (a *rpcError) Error() string {
	return ""
}

func (a *rpcError) GetError() error {
	return a.error
}

func (a *rpcError) GetCode() int {
	return a.code
}
