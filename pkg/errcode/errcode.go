package errcode

var (
	SUCCESS   = New(200, "success")
	ERROR     = New(500, "error")
	NOT_FOUND = New(404, "not found")
)

type ErrCode struct {
	code int
	msg  string
}

func New(code int, msg string) *ErrCode {
	return &ErrCode{
		code: code,
		msg:  msg,
	}
}

func (e *ErrCode) Code() int {
	return e.code
}

func (e *ErrCode) Msg() string {
	return e.msg
}
